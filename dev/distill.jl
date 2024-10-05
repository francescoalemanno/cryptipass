gorepr(c::Char) = c == '`' ? "\"`\"" : "`" * c * "`"

function putd(dd, from, to)
    if !(from in keys(dd))
        dd[from] = Dict{Char,Int}([])
    end
    if !(to in keys(dd[from]))
        dd[from][to] = 0
    end
    dd[from][to] += 1
end

function makesampler(io, rawV)
    sep = "   "
    S = collect(keys(rawV))
    C = collect(values(rawV))
    X = sortperm(-C)
    S = S[X]
    C = C[X]
    if length(S) == 1
        println(io, sep^2, "return ", gorepr(S[1]), ", 0.0")
        return
    end
    nor = sum(C)
    H = 0.0
    for c in C
        H -= log2(c / nor) * c / nor
    end
    println(io, sep^2, "H := ", H)
    println(io, sep^2, "r := g.Rng.IntN(", sum(C), ")")
    println(io, sep^2, "switch {")
    cum = 0
    for (s, c) in zip(S, C)
        println(io, sep^3, "case r < $(cum+c): return ", gorepr(s), ", H")
        cum += c
    end
    println(io, sep^3, "default: panic(\"unexpected rand num\")")
    println(io, sep^2, "}")
end

function whole_sampler(data)
    io = IOBuffer()
    kes = collect(keys(data))
    X = sortperm(map(length, values(data)))
    kes = kes[X]
    sep = "   "
    println(
        io,
        raw"""// PickNext returns the next character to the current seed string based on specific rules.
// It evaluates the last two characters of the seed to decide on the next character, either
// through predefined cases (such as 'mh' becoming 'mho') or randomly selecting a new character 
// with associated entropy.
//
// - Parameters:
//   seed (string): The current string being used to generate the passphrase.
//
// - Return values:
//   string: The updated string after appending the next character.
//   float64: The entropy contributed by the character selection process."""
    )
    println(io, "func (g *Generator) PickNext(seed string) (string, float64) {")
    println(io, sep, "L := min(len(seed),2)")
    println(io, sep, "tok := strings.ToLower(seed[len(seed)-L:])")
    println(io, sep, "retry:")
    println(io, sep, "switch tok {")
    for k in kes
        println(io, sep, "case `", k, "`:")
        makesampler(io, data[k])
    end
    println(io, sep, "default:")
    println(io, sep, sep, "tok=tok[1:]")
    println(io, sep, sep, "goto retry")
    println(io, sep, "}")
    println(io, "}")

    io
end


function makegensampler(rawV)
    io = IOBuffer()
    V = sort(collect(rawV))
    S = unique(V)
    C = [count(==(s), V) for s in S]
    X = sortperm(-C)
    S = S[X]
    C = C[X]
    nor = sum(C)
    H = 0.0
    for c in C
        H -= log2(c / nor) * c / nor
    end
    println(
        io,
        raw"""// PickLength returns a random word length and its associated entropy.
// The word length is chosen based on a weighted random selection,
// with lengths ranging from 3 to 9 characters.
//
// - Return values:
//   int: The selected word length in characters.
//   float64: The entropy contributed by the length selection process."""
    )
    cum = 0
    ifs = "if"
    println(io, "func (g *Generator) PickLength() (int, float64) {")
    println(io, "   ", "r := g.Rng.IntN(", sum(C), ")")
    println(io, "   ", "H := ", H)
    for (s, c) in zip(S, C)
        println(io, "   ", ifs, " r < ", cum + c, " {\n      return ", s, ", H")
        cum += c
        ifs = "} else if "
    end
    println(io, "   }\n   panic(\"unexpected rand num\")\n}")
    io
end


dict = map(readlines("./dev/eff.csv")) do w
    lowercase(replace(strip(w), '-' => ""))
end

alphabet = Set{Char}([])
for w in dict, c in w
    push!(alphabet, c)
end

lens = map(length, dict)

f22 = Dict{String,Dict{Char,Int}}()

for w in dict
    putd(f22, "", w[1])
    if length(w) > 1
        putd(f22, "" * w[1], w[2])
        I = collect(eachindex(w))
        for i in 1:length(I)-2
            key = w[I[i]] * w[I[i+1]]
            next = w[I[i+2]]
            putd(f22, key, next)
        end
    end
end





open("markovchain.go", "w") do io
    buf = whole_sampler(f22)
    write(io, "package cryptipass\n\nimport \"strings\"\n")
    write(io, "\n\n")
    write(io, "// THIS FILE HAS BEEN DISTILLED FROM EFF's long word list, without their work this software would not exist.")
    write(io, "\n\n")
    write(io, String(take!(buf)))
    write(io, "\n\n")
    write(io, String(take!(makegensampler(lens))))
end
run(`go fmt .`)

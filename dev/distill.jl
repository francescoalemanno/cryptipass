using DelimitedFiles, ProgressMeter

function simplify(vec)
    return vec
end

dict = map(readdlm("./dev/eff.csv", String)[:, 2]) do x
    replace(x, "-" => "")
end

lens = filter(<(8), sort(map(dict) do x
    length(x)
end))

for i = 1:470
    push!(lens, 8)
end

dget(d, k) = k in keys(d) ? d[k] : ""

f22 = Dict{String,String}()
enlarge_alphabet = true
if enlarge_alphabet
    consonant = "qwrtpsdfghjklzxcvbnm"
    vowels = "aeiouy"
    for x in 'a':'z'
        co = dget(f22, "")
        f22[""] = x in co ? co : co * x
    end

    for x in 'a':'z'
        key = "" * x
        if key in keys(f22)
            continue
        end
        f22[key] = (x in vowels ? consonant : vowels)
    end

    for c1 in 'a':'z', c2 in 'a':'z'
        key = c1 * c2
        if key in keys(f22)
            continue
        end
        f22[key] = (c2 in vowels ? consonant : vowels)
    end
end

@showprogress for w in dict
    for _ = 1:2
        f22[""] = collect(dget(f22, "") * w[1]) |> sort |> join
        f22[""*w[1]] = collect(dget(f22, "" * w[1]) * w[2]) |> sort |> join
        for i in 1:length(w)-2
            key = w[i] * w[i+1]
            next = w[i+2]
            past = dget(f22, key)
            f22[key] = collect(past * next) |> sort |> join
        end
    end
end




function makesampler(io, rawV)
    V = sort(collect(rawV))
    S = unique(V)
    C = [count(==(s), V) for s in S]
    nor = sum(C)
    H = 0.0
    for c in C
        H -= log2(c / nor) * c / nor
    end

    cum = 0
    ifs = "if"
    println(io, "   ", "r := rng.IntN(", sum(C), ")")
    println(io, "   ", "H := ", H)
    for (s, c) in zip(S, C)
        println(io, "   ", ifs, " ", cum, " <= r && r < ", cum + c, " {\n      return seed+", repr(string(s)), ", H")
        cum += c
        ifs = "} else if "
    end
    println(io, "   }\n   panic(\"unexpected rand num\")")
end

function whole_sampler(data)
    io = IOBuffer()
    kes = collect(keys(data))
    println(io, "\n\nfunc PickNext(seed string) (string, float64) {")
    println(io, "L := min(len(seed),2)")
    println(io, "tok := seed[len(seed)-L:]")
    for k in kes
        println(io, "if tok == ", repr(k), " {")
        makesampler(io, data[k])
        print(io, "} else ")
    end
    println(io, "{\n   panic(\"unexpected token \"+tok)\n}")
    println(io, "}")

    io
end


function makegensampler(rawV)
    io = IOBuffer()
    V = sort(collect(rawV))
    S = unique(V)
    C = [count(==(s), V) for s in S]

    nor = sum(C)
    H = 0.0
    for c in C
        H -= log2(c / nor) * c / nor
    end

    cum = 0
    ifs = "if"
    println(io, "func PickLength() (int, float64) {")
    println(io, "   ", "r := rng.IntN(", sum(C), ")")
    println(io, "   ", "H := ", H)
    for (s, c) in zip(S, C)
        println(io, "   ", ifs, " ", cum, " <= r && r < ", cum + c, " {\n      return ", s, ", H")
        cum += c
        ifs = "} else if "
    end
    println(io, "   }\n   panic(\"unexpected rand num\")\n}")
    io
end



open("markovchain.go", "w") do io
    buf = whole_sampler(f22)
    write(io, "package cryptipass\n")
    write(io, "\n\n")
    write(io, "// THIS FILE HAS BEEN DISTILLED FROM EFF's long word list, without their work this software would not exist.")
    write(io, "\n\n")
    write(io, String(take!(buf)))
    write(io, "\n\n")
    write(io, String(take!(makegensampler(lens))))
end

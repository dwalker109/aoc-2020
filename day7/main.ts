const input = await Deno.readTextFile("./input.txt");
const processed = input.split("\n").map(line => {
    const words = line.split(" ");
    const parent = words.slice(0, 2);
})

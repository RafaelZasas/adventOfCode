const std = @import("std");

var leftList: []i32 = undefined;
var rightList: []i32 = undefined;

const filePath = "input.txt";
// const filePath = "example.txt";

var buffer: [80000]u8 = undefined;
var fba = std.heap.FixedBufferAllocator.init(&buffer);
const allocator = fba.allocator();

fn init() !void {

    // Open the file in read mode
    const file = try std.fs.cwd().openFile(filePath, .{});
    defer file.close();

    // i guess this is a random pick for max buffer size
    const stat = try file.stat();
    const file_buffer = try file.readToEndAlloc(allocator, stat.size);

    var lines = std.ArrayList([]const u8).init(allocator);
    defer lines.deinit();

    var iter = std.mem.splitAny(u8, file_buffer, "\n");

    while (iter.next()) |line| {
        if (line.len == 0) {
            continue;
        }
        try lines.append(line);
    }

    const count = lines.items.len;

    // allocate enough mem for the arrays
    leftList = try allocator.alloc(i32, count);
    rightList = try allocator.alloc(i32, count);
    // defer allocator.free(leftList);
    // defer allocator.free(rightList);

    // iterate over the lines and split them into left and right
    var idx: usize = 0;

    for (0..count) |i| {
        const line = lines.items[i];
        var lineIter = std.mem.splitSequence(u8, line, "   ");

        const left = lineIter.next().?;
        const right = lineIter.next().?; // this is a bit weird

        leftList[idx] = try std.fmt.parseInt(i32, left, 10);
        rightList[idx] = try std.fmt.parseInt(i32, right, 10);

        idx += 1;
    }
}

fn partOne() void {
    var count: i32 = 0;

    // sort the arrays
    std.mem.sort(i32, leftList, {}, comptime std.sort.asc(i32));
    std.mem.sort(i32, rightList, {}, comptime std.sort.asc(i32));

    // iterate over leftList and rightList and calc distance
    for (0..leftList.len) |i| {
        const left = leftList[i];
        const right = rightList[i];

        const abs = getAbs(left, right);
        count += abs;
    }

    std.debug.print("partOne: {d}\n", .{count});
}

fn partTwo() !void {
    var similarityScore: i32 = 0;

    // create a dict to store the numbers and counts
    var countDict = std.AutoHashMap(i32, i32).init(allocator);
    defer countDict.deinit();

    for (0..rightList.len) |i| {

        // check if the abs is already in the dict
        const maybeCount = countDict.get(rightList[i]);

        if (maybeCount == null) {
            try countDict.put(rightList[i], 1);
        } else {
            try countDict.put(rightList[i], maybeCount.? + 1);
        }
    }

    for (0..leftList.len) |i| {
        const maybeCount = countDict.get(leftList[i]);

        var multiplier: i32 = 0;

        if (maybeCount != null) {
            multiplier = maybeCount.?;
        }

        similarityScore += leftList[i] * multiplier;
    }

    std.debug.print("partTwo: {d}\n", .{similarityScore});
}

fn getAbs(a: i32, b: i32) i32 {
    if (a > b) {
        return a - b;
    }
    return b - a;
}

pub fn main() !void {
    init() catch |err| {
        std.debug.print("Error: {any}\n", .{err});
    };

    var startTime: std.time.Timer = try std.time.Timer.start();
    partOne();
    // duration in nanoseconds
    const end = std.time.Timer.read(&startTime) / 1000;

    // print the duration in microseconds
    std.debug.print("Part One took: {d}µs\n", .{end});

    startTime = try std.time.Timer.start();
    partTwo() catch |err| {
        std.debug.print("Error: {any}\n", .{err});
    };
    const end2 = std.time.Timer.read(&startTime) / 1000;

    std.debug.print("Part Two took: {d}µs\n", .{end2});
}

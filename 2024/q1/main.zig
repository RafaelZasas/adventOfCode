const std = @import("std");

var leftList: []i32 = undefined;
var rightList: []i32 = undefined;

const filePath = "input.txt";
// const filePath = "example.txt";

fn init() !void {
    const allocator = std.heap.page_allocator;

    // Open the file in read mode
    const file = try std.fs.cwd().openFile(filePath, .{});
    defer file.close();

    // i guess this is a random pick for max buffer size
    const file_buffer = try file.readToEndAlloc(allocator, 32768);

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

    // start the timer
    var startTime: std.time.Timer = try std.time.Timer.start();
    partOne();
    // duration in nanoseconds
    const end = std.time.Timer.read(&startTime) / 1000;

    // print the duration in microseconds
    std.debug.print("Part One took: {d}µs\n", .{end});
}

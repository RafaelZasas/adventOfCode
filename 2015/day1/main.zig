const std = @import("std");

const allocator = std.heap.page_allocator;

const filePath = "input.txt";
// const filePath = "example.txt";

var fileBuffer: []const u8 = undefined;

fn init() !void {
    var startTime: std.time.Timer = try std.time.Timer.start();

    const file = try std.fs.cwd().openFile(filePath, .{});
    defer file.close();

    const stats = try file.stat();
    fileBuffer = try file.readToEndAlloc(allocator, stats.size);

    const end = std.time.Timer.read(&startTime) / 1000;
    std.debug.print("Init took: {d}µs\n", .{end});
}

fn partOne() !void {
    var startTime: std.time.Timer = try std.time.Timer.start();

    var floor: i16 = 0;

    for (fileBuffer) |c| {
        switch (c) {
            '(' => floor += 1,
            ')' => floor -= 1,
            else => {},
        }
    }
    const end = std.time.Timer.read(&startTime) / 1000;
    std.debug.print("Part one: {d}\n", .{floor});
    std.debug.print("Part one took: {d}µs\n", .{end});
}

fn partTwo() !void {
    var startTime: std.time.Timer = try std.time.Timer.start();

    var floor: i16 = 0;
    var position: usize = 0;

    for (0.., fileBuffer) |i, c| {
        if (c == '(') {
            floor += 1;
            if (floor == -1) {
                position = i + 1;
                break;
            }
            continue;
        }

        if (fileBuffer[i] == ')') {
            floor -= 1;
            if (floor == -1) {
                position = i + 1;
                break;
            }
        }
    }

    const end = std.time.Timer.read(&startTime) / 1000;
    std.debug.print("Part two: {d}\n", .{position});
    std.debug.print("Part two took: {d}µs\n", .{end});
}

pub fn main() !void {
    try init();
    try partOne();
    try partTwo();
    allocator.free(fileBuffer);
}

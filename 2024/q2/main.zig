const std = @import("std");

const allocator = std.heap.page_allocator;
var matrix: std.ArrayList([]const u8) = undefined;

fn init() !void {
    var startTime: std.time.Timer = try std.time.Timer.start();

    matrix = std.ArrayList([][]const u8).init(allocator);

    const file = try std.fs.cwd().openFile("input.txt", .{});
    defer file.close();

    const stats = try file.stat();
    const file_buffer = try file.readToEndAlloc(allocator, stats.size);

    var lines = std.ArrayList([]const u8).init(allocator);
    defer lines.deinit();

    var iterator = std.mem.splitAny(u8, file_buffer, "\n");
    while (iterator.next()) |line| {
        if (line.len == 0) {
            continue;
        }
        try lines.append(line);
    }

    var idx: usize = 0;
    for (lines.items) |line| {
        try matrix.append(line);
        // var lineIter = std.mem.splitSequence(u8, line, " ");
        //
        // while (lineIter.next()) |number| {
        //     if (number.len == 0) {
        //         continue;
        //     }
        //     const num = try std.fmt.parseInt(u8, number, 10);
        //     try matrix.items[idx].append(num);
        // }

        idx += 1;
    }

    // Access the 2D array
    for (0..matrix.len) |i| {
        for (0..matrix[i].len) |j| {
            // convert the i32 to string
            std.debug.print("{d} ", .{matrix[i][j]});
        }
        std.debug.print("\n", .{});
    }
    const end = std.time.Timer.read(&startTime) / 1000;
    std.debug.print("Init took: {d}µs\n", .{end});
}

pub fn main() !void {
    defer matrix.deinit();
    try init();
}

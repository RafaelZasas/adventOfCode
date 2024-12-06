const std = @import("std");

const allocator = std.heap.page_allocator;

const filePath = "input.txt";
// const filePath = "example.txt";

var fileBuffer: []const u8 = undefined;
var volumes: []i32 = undefined;
var perimeters: []i32 = undefined;

fn init() !void {
    var startTime: std.time.Timer = try std.time.Timer.start();

    const file = try std.fs.cwd().openFile(filePath, .{});
    defer file.close();

    const stats = try file.stat();
    fileBuffer = try file.readToEndAlloc(allocator, stats.size);

    const end = std.time.Timer.read(&startTime) / 1000;
    std.debug.print("Init took: {d}µs\n", .{end});
}

// find surface area
// 2(lw + lh + wh) + area of smallest side
fn partOne() !void {
    var startTime: std.time.Timer = try std.time.Timer.start();

    // l,w,h separated by x
    var totalArea: i32 = 0;

    var lines = std.ArrayList([]const u8).init(allocator);
    defer lines.deinit();

    var iterator = std.mem.splitAny(u8, fileBuffer, "\n");
    while (iterator.next()) |line| {
        if (line.len == 0) {
            continue;
        }
        try lines.append(line);
    }

    perimeters = try allocator.alloc(i32, lines.items.len);
    volumes = try allocator.alloc(i32, lines.items.len);

    for (lines.items, 0..) |line, idx| {
        var xIter = std.mem.splitAny(u8, line, "x");
        const l = try std.fmt.parseInt(i32, xIter.next().?, 10);
        const w = try std.fmt.parseInt(i32, xIter.next().?, 10);
        const h = try std.fmt.parseInt(i32, xIter.next().?, 10);

        const lw = l * w;
        const lh = l * h;
        const wh = w * h;

        const area = 2 * (lw + lh + wh);
        var smallest: i32 = 0;
        if (lw < lh) {
            if (lw < wh) {
                smallest = lw;
            } else {
                smallest = wh;
            }
        } else {
            if (lh < wh) {
                smallest = lh;
            } else {
                smallest = wh;
            }
        }

        var perimeterOpts: [3]i32 = undefined;

        perimeterOpts[0] = 2 * (l + w);
        perimeterOpts[1] = 2 * (l + h);
        perimeterOpts[2] = 2 * (w + h);

        var smallestPerimeter: i32 = perimeterOpts[0];
        for (perimeterOpts) |perimeter| {
            if (perimeter < smallestPerimeter) {
                smallestPerimeter = perimeter;
            }
        }

        perimeters[idx] = smallestPerimeter;
        volumes[idx] = l * w * h;

        totalArea += area + smallest;
    }

    const end = std.time.Timer.read(&startTime) / 1000;
    std.debug.print("Part one: {d}\n", .{totalArea});
    std.debug.print("Part one took: {d}µs\n", .{end});
}

// get ribbon length from smallest perimeter of any face
// and bow length from volume
fn partTwo() !void {
    var startTime: std.time.Timer = try std.time.Timer.start();

    var totalLength: i32 = 0;

    for (perimeters, 0..) |ribbon, idx| {
        const bow = volumes[idx];

        totalLength += ribbon + bow;
    }

    const end = std.time.Timer.read(&startTime) / 1000;
    std.debug.print("Part two: {d}\n", .{totalLength});
    std.debug.print("Part two took: {d}µs\n", .{end});
}

pub fn main() !void {
    try init();
    try partOne();
    try partTwo();
    allocator.free(fileBuffer);
}

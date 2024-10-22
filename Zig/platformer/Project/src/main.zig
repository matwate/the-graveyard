const rl = @import("raylib");

pub fn main() anyerror!void {
    const window_width: u64 = 1600;
    const window_height: u64 = 900;

    rl.initWindow(window_width, window_height, "Platformer testing")
    defer rl.closeWindow();

    while(!rl.windowShouldClose()) {
        rl.beginDrawing();
        defer rl.endDrawing();

        rl.clearBackground(rl.Color.White);
        rl.drawText("Window")
    }

    

}

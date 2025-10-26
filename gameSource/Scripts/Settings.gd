extends Control
@onready var fullscreen: Button = $fullscreen
@onready var vsync: Button = $VSYNC

var is_fullscreen: bool = false
var is_vsync: bool = true

func _on_fullscreen_pressed() -> void:
	if not is_fullscreen:
		fullscreen.text = "DISABLE FULLSCREEN"
		DisplayServer.window_set_mode(DisplayServer.WINDOW_MODE_FULLSCREEN)
	else:
		fullscreen.text = "ENABLE FULLSCREEN"
		DisplayServer.window_set_mode(DisplayServer.WINDOW_MODE_WINDOWED)
	is_fullscreen = not is_fullscreen


func _on_vsync_pressed() -> void:
	if not is_vsync:
		vsync.text = "DISABLE VSYNC"
		DisplayServer.window_set_vsync_mode(DisplayServer.VSYNC_ENABLED)
	else:
		vsync.text = "ENABLE VSYNC"
		DisplayServer.window_set_vsync_mode(DisplayServer.VSYNC_DISABLED)
	is_vsync = not is_vsync

func _on_return_pressed() -> void:
	get_tree().change_scene_to_file("res://Scenes/MENU.tscn")

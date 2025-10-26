extends Control
@onready var main: Control = $Main

func _on_start_pressed() -> void:
	get_tree().change_scene_to_file("res://Scenes/Game.tscn")

func _on_options_pressed() -> void:
	get_tree().change_scene_to_file("res://Scenes/Settings.tscn")

func _on_exit_pressed() -> void:
	get_tree().quit()

func _on_game_2_pressed() -> void:
	get_tree().change_scene_to_file("res://Scenes/Game2.tscn")

extends Node2D
@onready var player: Node2D = $"."

var is_dead: bool = false

func _process(_delta: float) -> void:
	if not is_dead:
		var pos = get_global_mouse_position().x
		player.global_position.x = pos - 50

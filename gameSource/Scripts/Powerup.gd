extends Node2D

signal collected

func _ready() -> void:
	position.x = randi_range(0,1100)

func _process(_delta: float) -> void:
	position.y += 4

func _on_area_2d_area_entered(area: Area2D) -> void:
	if area.name == "Player":
		collected.emit()
		call_deferred("queue_free")

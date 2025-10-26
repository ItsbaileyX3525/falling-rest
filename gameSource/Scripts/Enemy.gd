extends Node2D

@onready var area_2d: Area2D = $Area2D
@onready var texture_rect: TextureRect = $TextureRect

signal collected(value: int)
signal missed(value: int)

const LEAF_1 = preload("uid://dkgoup7rdlkkq") #Orange leaf - Normal
const LEAF_2 = preload("uid://bw6i7ncmwvfek") #Yellow leaf - Fast/Heavy
const LEAF_3 = preload("uid://dpu7ap22fqphj") #Red leaf - Worth 2 leaf - Slower

var leaf_type: int = 0
var drops_slowed: bool = false
var drops_speed: bool = false
var drop_multi: float = 1.0

func _init() -> void:
	position.x = randf_range(0,1100)

func _ready() -> void:
	leaf_type = randi_range(1,3)
	if drops_slowed:
		drop_multi = 0.2
	if drops_speed: #Curse takes power
		drop_multi = 1.5
	match leaf_type:
		1:
			print("Orange leaf")
			texture_rect.texture = LEAF_1
		2:
			print("Yellow leaf")
			texture_rect.texture = LEAF_2
		3:
			print("Red leaf")
			texture_rect.texture = LEAF_3


func _physics_process(_delta: float) -> void:
	position.y += 5
	match leaf_type:
		1:
			position.y += (5 * drop_multi)
		2:
			position.y += (8 * drop_multi)
		3:
			position.y += (2 * drop_multi)

func _on_area_2d_area_entered(area: Area2D) -> void:
	if area.name == "Player":
		match leaf_type:
			1:
				collected.emit(1)
			2:
				collected.emit(1)
			3:
				collected.emit(2)
		call_deferred("queue_free")
	elif area.name == "Death":
		match leaf_type:
			1:
				missed.emit(1)
			2:
				missed.emit(1)
			3:
				missed.emit(2)
		call_deferred("queue_free")

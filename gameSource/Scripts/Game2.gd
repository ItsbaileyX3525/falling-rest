extends Node2D
@onready var score_label: Label = $MainArea/Score
@onready var leaves_area: Control = $LeavesArea
@onready var click_power: Button = $MainArea/VBoxContainer/ClickPower
@onready var click_power_2: Button = $MainArea/VBoxContainer/ClickPower2
@onready var generator: Timer = $Generator
@onready var gen_speeds: Button = $MainArea/VBoxContainer/GenSpeeds
@onready var gen_speeds_2: Button = $MainArea/VBoxContainer/GenSpeeds2

const LEAF_1 = preload("uid://dkgoup7rdlkkq")
const LEAF_2 = preload("uid://bw6i7ncmwvfek")
const LEAF_3 = preload("uid://dpu7ap22fqphj")

var score: float = 1.0
var click_multi: float = 1.0
var gen1_price: int = 75
var gen2_price: int = 200
var gen3_price: int = 5000
var gen4_price: int = 25000
var gen5_price: int = 50000
var gen6_price: int = 75000
var timerWaitTime: float = 1.0

var generating_leaves: int = 0

func _process(_delta: float) -> void:
	score_label.text = "Leaves: %s" % score

func _on_clicker_pressed() -> void:
	score += (1 * click_multi)
	var randomLeaf = randi_range(1,3)
	var textureRect = TextureRect.new()
	match randomLeaf:
		1:
			textureRect.texture = LEAF_1
		2:
			textureRect.texture = LEAF_2
		3:
			textureRect.texture = LEAF_3
	
	leaves_area.add_child(textureRect)
	textureRect.position = Vector2(randi_range(0,1100),randi_range(0,600))
	textureRect.scale = Vector2(0.165,0.165)
	await get_tree().create_timer(1).timeout
	
	textureRect.call_deferred("queue_free")

func _on_click_power_pressed() -> void:
	if score >= 100:
		score -= 100
		click_multi += 1
		click_power.call_deferred("queue_free")

func _on_click_power_2_pressed() -> void:
	if score >= 2500:
		score -= 2500
		click_multi += 3
		click_power_2.call_deferred("queue_free")

func _on_gen_1_pressed() -> void:
	if score >= gen1_price:
		score -= gen1_price
		var newPrice = int(floor(gen1_price * 1.25))
		gen1_price = newPrice
		generating_leaves += 2

func _on_gen_2_pressed() -> void:
	if score >= gen2_price:
		score -= gen2_price
		var newPrice = int(floor(gen2_price * 1.25))
		gen2_price = newPrice
		generating_leaves += 5

func _on_gen_3_pressed() -> void:
	if score >= gen3_price:
		score -= gen3_price
		var newPrice = int(floor(gen3_price * 1.25))
		gen3_price = newPrice
		generating_leaves += 15

func _on_gen_4_pressed() -> void:
	if score >= gen4_price:
		score -= gen4_price
		var newPrice = int(floor(gen4_price * 1.25))
		gen4_price = newPrice
		generating_leaves += 50

func _on_gen_5_pressed() -> void:
	if score >= gen5_price:
		score -= gen5_price
		var newPrice = int(floor(gen5_price * 1.25))
		gen5_price = newPrice
		generating_leaves += 500

func _on_gen_6_pressed() -> void:
	if score >= gen6_price:
		score -= gen6_price
		var newPrice = int(floor(gen6_price * 1.25))
		gen6_price = newPrice
		generating_leaves += 6700

func _on_generator_timeout() -> void:
	score += generating_leaves

func _on_gen_speeds_pressed() -> void:
	if score >= 800:
		score -= 800
		timerWaitTime -= .2
		generator.wait_time = timerWaitTime
		gen_speeds.call_deferred("queue_free")

func _on_gen_speeds_2_pressed() -> void:
	if score >= 5000:
		score -= 5000
		timerWaitTime -= .5
		generator.wait_time = timerWaitTime
		gen_speeds_2.call_deferred("queue_free")

func _on_menu_pressed() -> void:
	get_tree().change_scene_to_file("res://Scenes/MENU.tscn")

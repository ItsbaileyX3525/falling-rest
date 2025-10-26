extends Node2D

@onready var player: Node2D = $Player
@onready var drops: Node2D = $Drops
@onready var enemy_spawn_timer: Timer = $EnemySpawnTimer
@onready var score_text: Label = $UI/Score
@onready var lives_text: Label = $UI/Lives
@onready var powerup_spawn_timer: Timer = $PowerupSpawnTimer
@onready var animation_player: AnimationPlayer = $UI/tomato/AnimationPlayer
@onready var tomato: Sprite2D = $UI/tomato
@onready var game_over_screen: Control = $UI/GameOverScreen
@onready var powerup_text: Label = $UI/Powerup
@onready var hide_powerup: Timer = $HidePowerup

const ENEMY = preload("res://Scenes/Enemy.tscn")
const POWERUP = preload("uid://dylka5kh87rkx")

var lives: int = 3
var score: int = 0
var timer: float = 4.0
var drops_slowed: bool = false
var drops_speed: bool = false

func _process(_delta: float) -> void:
	score_text.text = "SCORE: %s" % score
	lives_text.text = "LIVES: %s" % lives

func _ready() -> void:
	enemy_spawn_timer.start()
	powerup_spawn_timer.start()

func _on_enemy_collected(value: int) -> void:
	score += value

func _on_enemy_missed(value: int) -> void:
	lives -= value
	if lives <= 0:
		lives = 0
		enemy_spawn_timer.stop()
		powerup_spawn_timer.stop()
		tomato.visible = true
		player.is_dead = true
		animation_player.play("throw")

func _on_enemy_spawn_timer_timeout() -> void:
	var enemy_instance = ENEMY.instantiate()
	enemy_instance.connect("collected", _on_enemy_collected)
	enemy_instance.connect("missed", _on_enemy_missed)
	enemy_instance.drops_slowed = drops_slowed
	enemy_instance.drops_speed = drops_speed
	drops.add_child(enemy_instance)
	if timer >= .7:
		timer -= .12
	enemy_spawn_timer.wait_time = timer
	enemy_spawn_timer.start()

func powerup_collected() -> void:
	var powerupType = randi_range(1,5)
	powerup_text.visible = true
	hide_powerup.start()
	match powerupType:
		1:
			powerup_text.text = "Extra lives"
			lives += 2
			if lives > 10: #Hard cap 10
				lives = 10
		2:
			powerup_text.text = "Bonus points"
			score += randi_range(1,3)
		3:
			powerup_text.text = "Let's slow it down!"
			if drops_slowed:
				return
			drops_slowed = true
			await get_tree().create_timer(5).timeout
			drops_slowed = false
		4:
			powerup_text.text = "HAHA LOST POINTS"
			score -= randi_range(1,5)
			if score < 0:
				score = 0
		5:
			powerup_text.text = "I'm feeling a little speedy"
			if drops_speed:
				return
			drops_speed = true
			await get_tree().create_timer(5).timeout
			drops_speed = false

func _on_powerup_spawn_timer_timeout() -> void:
	var powerupInstance = POWERUP.instantiate()
	var timerPowerup: float

	powerupInstance.connect("collected", powerup_collected)
	drops.add_child(powerupInstance)
	timerPowerup = randf_range(3,10)
	powerup_spawn_timer.wait_time = timerPowerup
	powerup_spawn_timer.start()

func _on_animation_player_animation_finished(anim_name: StringName) -> void:
	if anim_name == "throw":
		game_over_screen.visible = true

func _on_hide_powerup_timeout() -> void:
	powerup_text.visible = false


func _on_menu_pressed() -> void:
	get_tree().change_scene_to_file("res://Scenes/MENU.tscn")

func _on_quit_pressed() -> void:
	get_tree().quit()

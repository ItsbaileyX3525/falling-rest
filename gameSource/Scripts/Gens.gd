extends Button

@onready var gens_area: Control = $"../GensArea"
@onready var main_area: Control = $"../MainArea"


func _on_pressed() -> void:
	if main_area.visible:
		gens_area.visible = true
		main_area.visible = false
	else:
		gens_area.visible = false
		main_area.visible = true

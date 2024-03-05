local state = require("State")
local Character = require("Character")
local Engine = {
	State = state,
}

function Engine:InitDefault()
	state:InitDefault()
end
function Engine:AddPlayer(player)
	state:AddPlayer(player)
end
function Engine:DrawActivityCards(n)
	state:DrawActivityCards(n)
end
function Engine:StartPlayerTurn()
	state:StartPlayerTurn()
end
function Engine:PlayPlayerCard(playerIndex, cardIndex)
	player = state.GetPlayer(playerIndex)
end
return Engine


local Card = require "Card"
local repo = require "Repository"
local State = require "State"
local e = require "Engine"
local r = require "Renderer"
local Character = require "Character"

function love.load()
    love.window.setTitle("Atipical House")
    e:InitDefault()
    e:AddPlayer(Character:New{
        Name = "Orlando",
        Character = "Bloom",
    })
    e:DrawActivityCards(5)
    e:StartPlayerTurn()
end

function love.draw()
    -- working 
    -- local d = Card:ToString()
    -- local d = Card:New{Title = "Hola"}:ToString()

   -- local d = repo:GetActivity(0):ToString()
   -- love.graphics.print(d, 0, 100)

    r.DrawEngineState()
  --  for i = 1, 3 do
  --      State:DrawActivityCards()
  --  end
end

function love.mousepressed(x, y, button, istouch)
    if button == 1 then
       e:DrawActivityCards(1)
       e:StartPlayerTurn()
    end
 end
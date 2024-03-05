local e = require "Engine"
local Renderer = {}

local mx, my

function Renderer:DrawEngineState()

    mx, my = love.mouse.getPosition()
  --  love.graphics.print("this is the renderer", 10, 10)
    love.graphics.print(e.State:ActivityPileSize(), 300, 10)
    Renderer:DrawBoard()
    Renderer:DrawPlayer(e.State.Players[1])
end 
function Renderer:DrawBoard()
   -- loop through the active cards
    for i = 1, #e.State.board.activeCards do
        local card = e.State.board.activeCards[i]

        -- if mouse is over the card
        if mx > 300 and mx < 300 + 300 and my > 20 * i and my < 20 * i + 19 then
            love.graphics.setColor(255, 0, 100, 255)
            love.graphics.rectangle("fill", 300, 20*i, 300, 19)
            love.graphics.setColor(255, 255, 255, 255)
            love.graphics.print(card:ToString(), 300, 20 * i)
        else
            love.graphics.setColor(255, 100, 255, 255)
            love.graphics.rectangle("fill", 300, 20*i, 300, 19)
            love.graphics.setColor(0, 0, 0, 255)
            love.graphics.print(card:ToString(), 300, 20 * i)
        end
        -- love.graphics.rectangle("fill", 300, 20*i, 300, 19)

        -- love.graphics.setColor(255, 255, 255, 255)
        -- love.graphics.print(card:ToString(), 300, 20 * i)
    end

  
end

function Renderer:DrawPlayer(p)
    local x, y = 10, 10
    love.graphics.setColor(255, 0, 100, 255)
    love.graphics.print(p.Name .. " " .. p.Energy .. "/" .. p.Capacity .. "/" .. p.MaxCapacity, x, y)
    y = y + 10
    love.graphics.print(" --- Hand", x, y)
    y = y + 10
    if(p.Deck.hand == nil) then
        love.graphics.print(" Empty ", x, y)    
    else
        --love.graphics.print(p.Deck.hand[1]:ToString(), x, y)
        for i = 1, #p.Deck.hand do
            -- if mouse over the card 
            local card = p.Deck.hand[i]
            if mx > x and mx < x + 300 and my > 20 * i + y and my < 20 * i +y + 19 then                
                love.graphics.setColor(255, 0, 100, 255)
                love.graphics.rectangle("fill", x, 20*i + y, 300, 19)
                love.graphics.setColor(255, 255, 255, 255)
                love.graphics.print(card:ToString(), x, 20 * i + y)              
            else
                love.graphics.setColor(255, 100, 255, 255)
                love.graphics.rectangle("fill", x, 20*i +y, 300, 19)
                love.graphics.setColor(0, 0, 0, 255)
                love.graphics.print(card:ToString(), x, 20 * i + y)
            end 
        end
    end
    
end

return Renderer
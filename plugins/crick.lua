--
-- Created by IntelliJ IDEA.
-- User: NightWolf
-- Date: 12/16/2017
-- Time: 12:05 AM
-- To change this template use File | Settings | File Templates.
--

--[[ Loaded Libraries ]]--
http = require("http")
json = require("json")

--[[ Global Variables ]]--
user = nil

--[[ Helper function for creating an array of frames from a table ]]--
function create_frame_array(raw_frames)
    frames = {}
    for k, v in pairs(raw_frames) do
        frames[k] = frame.new(v['ID'], v['Name'])
    end
    return frames
end

function init()
    --[[
    response, error_message = http.get("https://jsonplaceholder.typicode.com/users")
    users = json.decode(response.body)
    for _, v in pairs(users) do
        print(k)
        for k_2, v_2 in pairs(v) do
            print("\t" .. k_2)
            if ( type(v_2) == "table")  then
                for k_3, v_3 in pairs(v_2) do
                    if ( type(v_3) == "table") then
                        print("\t\tIs a table!")
                    else
                        print("\t\t" .. k_3 .. " : " .. v_3)
                    end
                end
            else
                print("\t\t" .. v_2)
            end
        end
    end
    ]]--

    f = frame.new("1", "Hello World")
end

--[[ Sync frames. Should return a list of frames ]]--
function sync(raw_frames)
    raw_frames = json.decode(raw_frames)
    frames = create_frame_array(raw_frames)
    print(frames[1]:name())
    return raw_frames
end

function cleanup(config)

end


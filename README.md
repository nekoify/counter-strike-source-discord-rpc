# Counter Strike: Source Discord RPC

### ONLY WORKS ON WINDOWS 
It should be VAC safe since it only reads memory but its not my fault if anything happens :p
## Running automatically on CS:S launch
Start by creating a batch file with the following contents
```bat
@echo off
echo Starting Counter-Strike: Source...
start "" %1 -steam
echo Waiting for 20 seconds before starting Discord RPC...
timeout /t 20 /nobreak >nul
echo Starting Discord RPC...
start "" /B cmd /c "C:\Path\To\counter-strike-source-discord-rpc.exe"
exit
```
 Next go to Steam and find CS:S, right click it, go to properties and find launch options
 <br>
 Paste the following
 <br>
 `C:\Path\To\batName.bat %command%
 `
<br/>

## Screenshot

![Screenshot](https://raw.githubusercontent.com/nekoify/counter-strike-source-discord-rpc/main/screenshots/screenshot.png)

# Step 0: Install the Twilio CLI and authenticated with account API Credentials or with a Master API Key
# Twilio CLI installation instructions: https://www.twilio.com/docs/twilio-cli/quickstart
function WriteLog
{
Param ([string]$LogString)

$DateTime = "[{0:MM/dd/yy} {0:HH:mm:ss}]" -f (Get-Date)
$LogMessage = "$Datetime $LogString"

Write-Output $LogMessage
}


# Step 1: Create a new API Key
$NewApiKey = twilio api:core:keys:create -o json | ConvertFrom-Json;
$NewApiKeySid = $NewApiKey.sid;
$NewApiKeySecret = $NewApiKey.secret;

WriteLog  "ApiSidKey=$NewApiKeySid"
WriteLog "ApiSecret=$NewApiKeySecret"
# Step 2: Update your applications to use the new API Key SID and API Key Secret
# --- TO IMPLEMENT BY YOU --- 

# Step 3: Fetch the existing API Key SID (hardcoded for sample)
$OriginalApiKeySid = [System.Environment]::GetEnvironmentVariable('TWILIO_SID_DREWS_MAIN', 'User');
# Step 4: Delete the old API Key
twilio api:core:keys:remove --sid=$OriginalApiKeySid;

WriteLog "removed old sid"
# Step 5: set the new api key sid
[Environment]::SetEnvironmentVariable("TWILIO_SID_DREWS_MAIN",$NewApiKeySid,"User")
WriteLog "wrote the environment variable for twilio SID:$NewApiKeySid"
# Step 6: set the new api key secret
[Environment]::SetEnvironmentVariable("TWILIO_SECRET_DREWS_MAIN",$NewApiKeySecret,"User")
WriteLog "wrote the environment variable for twilio Secret:$NewApiKeySecret"

WriteLog "exiting..."

Exit-PSSession
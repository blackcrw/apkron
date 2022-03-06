package interesting

import (
	"fmt"
	"os"
	"regexp"
)

func AndroidManifest(temporary_directory string) ([][]string, error) {
	var _ = [78][2]string{
		{"ACCESS_CHECKIN_PROPERTIES"            , "Allows read/write access to the \"properties\" table in the checkin database, to change values that get uploaded."},
		{"ACCESS_COARSE_LOCATION"               , "Allows an app to access approximate location derived from network location sources such as cell towers and Wi-Fi."},
		{"ACCESS_FINE_LOCATION"                 , "Allows an app to access precise location from location sources such as GPS, cell towers, and Wi-Fi."},
		{"ACCESS_LOCATION_EXTRA_COMMANDS"       , "Allows an application to access extra location provider commands."},		
		{"ACCESS_MOCK_LOCATION"                 , "Allows an application to create mock location providers for testing."},
		{"ACCESS_NETWORK_STATE"                 , "Allows applications to access information about networks."},
		{"ACCESS_NOTIFICATION_POLICY"           , "Marker permission for applications that wish to access notification policy."},
		{"ACCESS_SURFACE_FLINGER"               , "Allows an application to use SurfaceFlinger's low level features."},
		{"ACCESS_WIFI_STATE"                    , "Allows applications to access information about Wi-Fi networks"},
		{"ACCOUNT_MANAGER"                      , "Allows applications to call into AccountAuthenticators."},		
		{"ADD_VOICEMAIL"                        , "Allows an application to add voicemails into the system."},
		{"AUTHENTICATE_ACCOUNTS"                , "Allows an application to act as an AccountAuthenticator for the AccountManager."},
		{"BATTERY_STATS"                        , "Allows an application to collect battery statistics."},
		{"BIND_ACCESSIBILITY_SERVICE"           , "Must be required by an AccessibilityService,to ensure that only the system can bind to it."},
		{"BIND_APPWIDGET"                       , "Allows an application to tell the AppWidget service which application can access AppWidget\"s data."},
		{"BIND_CARRIER_MESSAGING_SERVICE"       , "The system process that is allowed to bind to services in carrier apps will have this permission. Carrier apps should use this permission to protect their services that only the system is allowed to bind to."},
		{"BIND_CHOOSER_TARGET_SERVICE"          , "Must be required by a ChooserTargetService, to ensure that only the system can bind to it."},
		{"BIND_DEVICE_ADMIN"                    , "Must be required by device administration receiver, to ensure that only the system can interact with it."},
		{"BIND_DREAM_SERVICE"                   , "Must be required by an DreamService, to ensure that only the system can bind to it."},
		{"BIND_INCALL_SERVICE"                  , "Must be required by a InCallService, to ensure that only the system can bind to it."},
		{"BIND_INPUT_METHOD"                    , "Must be required by an InputMethodService, to ensure that only the system can bind to it."},
		{"BIND_MIDI_DEVICE_SERVICE"             , "Must be required by an MidiDeviceService, to ensure that only the system can bind to it."},
		{"BIND_NFC_SERVICE"                     , "Must be required by a HostApduService or OffHostApduService to ensure that only the system can bind to it."},
		{"BIND_NOTIFICATION_LISTENER_SERVICE"   , "Must be required by an NotificationListenerService, to ensure that only the system can bind to it."},
		{"BIND_PRINT_SERVICE"                   , "Must be required by a PrintService, to ensure that only the system can bind to it."},
		{"BIND_REMOTEVIEWS"                     , "Must be required by a RemoteViewsService, to ensure that only the system can bind to it."},
		{"BIND_TELECOM_CONNECTION_SERVICE"      , "Must be required by a ConnectionService, to ensure that only the system can bind to it."},
		{"BIND_TEXT_SERVICE"                    , "Must be required by a TextService (e.g. SpellCheckerService) to ensure that only the system can bind to it."},
		{"BIND_TV_INPUT"                        , "Must be required by a TvInputService to ensure that only the system can bind to it."},		
		{"BIND_VOICE_INTERACTION"               , "Must be required by a VoiceInteractionService, to ensure that only the system can bind to it."},
		{"BIND_VPN_SERVICE"                     , "Must be required by a VpnService, to ensure that only the system can bind to it."},
		{"BIND_WALLPAPER"                       , "Must be required by a WallpaperService, to ensure that only the system can bind to it."},
		{"BLUETOOTH"                            , "Allows applications to connect to paired bluetooth devices."},
		{"BLUETOOTH_ADMIN"                      , "Allows applications to discover and pair bluetooth devices."},
		{"BLUETOOTH_PRIVILEGED"                 , "Allows applications to pair bluetooth devices without user interaction."},
		{"BODY_SENSORS"                         , "Allows an application to access data from sensors that the user uses to measure what is happening inside his/her body, such as heart rate."},
		{"BRICK"                                , "Required to be able to disable the device (very dangerous!)."},
		{"BROADCAST_PACKAGE_REMOVED"            , "Allows an application to broadcast a notification that an application package has been removed."},
		{"BROADCAST_SMS"                        , "Allows an application to broadcast an SMS receipt notification."},
		{"BROADCAST_STICKY"                     , "Allows an application to broadcast sticky intents."},
		{"BROADCAST_WAP_PUSH"                   , "Allows an application to broadcast a WAP PUSH receipt notification."},
		{"CALL_PHONE"                           , "Allows an application to initiate a phone call without going through the Dialer user interface for the user to confirm the call being placed."},
		{"CALL_PRIVILEGED"                      , "Allows an application to call any phone number, including emergency numbers, without going through the Dialer user interface for the user to confirm the call being placed."},
		{"CAMERA"                               , "Required to be able to access the camera device."},
		{"CAPTURE_AUDIO_OUTPUT"                 , "Allows an application to capture audio output."},
		{"CAPTURE_SECURE_VIDEO_OUTPUT"          , "Allows an application to capture secure video output."},		
		{"CAPTURE_VIDEO_OUTPUT"                 , "Allows an application to capture video output."},
		{"CHANGE_COMPONENT_ENABLED_STATE"       , "Allows an application to change whether an application component (other than its own) is enabled or not."},
		{"CHANGE_CONFIGURATION"                 , "Allows an application to modify the current configuration, such as locale."},
		{"CHANGE_NETWORK_STATE"                 , "Allows applications to change network connectivity state."},
		{"CHANGE_WIFI_MULTICAST_STATE"          , "Allows applications to enter Wi-Fi Multicast mode."},
		{"CHANGE_WIFI_STATE"                    , "Allows applications to change Wi-Fi connectivity state."},
		{"CLEAR_APP_CACHE"                      , "Allows an application to clear the caches of all installed applications on the device."},
		{"CLEAR_APP_USER_DATA"                  , "Allows an application to clear user data."},
		{"CONTROL_LOCATION_UPDATES"             , "Allows enabling/disabling location update notifications from the radio."},
		{"DELETE_CACHE_FILES"                   , "Allows an application to delete cache files."},
		{"DELETE_PACKAGES"                      , "Allows an application to delete packages."},
		{"DEVICE_POWER"                         , "Allows low-level access to power management."},
		{"DIAGNOSTIC"                           , "Allows applications to RW to diagnostic resources."},
		{"DISABLE_KEYGUARD"                     , "Allows applications to disable the keyguard."},
		{"DUMP"                                 , "Allows an application to retrieve state dump information from system services."},		
		{"EXPAND_STATUS_BAR"                    , "Allows an application to expand or collapse the status bar."},
		{"FACTORY_TEST"                         , "Run as a manufacturer test application, running as the root user."},
		{"FLASHLIGHT"                           , "Allows access to the flashlight."},
		{"FORCE_BACK"                           , "Allows an application to force a BACK operation on whatever is the top activity."},
		{"GET_ACCOUNTS"                         , "Allows access to the list of accounts in the Accounts Service."},
		{"GET_ACCOUNTS_PRIVILEGED"              , "Allows access to the list of accounts in the Accounts Service."},
		{"GET_PACKAGE_SIZE"                     , "Allows an application to find out the space used by any package."},
		{"GET_TASKS"                            , "Allows an application to get information about the currently or recently running tasks."},
		{"GET_TOP_ACTIVITY_INFO"                , "Allows an application to retrieve private information about the current top activity, such as any assist context it can provide."},
		{"GLOBAL_SEARCH"                        , "This permission can be used on content providers to allow the global search system to access their data."},
		{"HARDWARE_TEST"                        , "Allows access to hardware peripherals."},
		{"INJECT_EVENTS"                        , "Allows an application to inject user events (keys, touch, trackball) into the event stream and deliver them to ANY window."},
		{"INSTALL_LOCATION_PROVIDER"            , "Allows an application to install a location provider into the Location Manager."},
		{"INSTALL_PACKAGES"                     , "Allows an application to install packages."},
		{"INSTALL_SHORTCUT"                     , "Allows an application to install a shortcut in Launcher."},
		{"INTERNAL_SYSTEM_WINDOW"               , "Allows an application to open windows that are for use by parts of the system user interface."},
		{"INTERNET"                             , "Allows applications to open network sockets."},
	}

	var raw, err = os.ReadFile(temporary_directory+"/AndroidManifest.xml")

	if err != nil { return [][]string{}, fmt.Errorf("%w", err) }
		
	var regxp, _ = regexp.Compile(`(?i)android:permission=\"android.permission.(.*?)\".*?`)

	var submatch = regxp.FindAllStringSubmatch(string(raw[:]), -1)
	
	return submatch, nil
}
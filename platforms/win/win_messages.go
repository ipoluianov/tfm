package win

/*
#define WM_NULL                         0x0000
#define WM_CREATE                       0x0001
#define WM_DESTROY                      0x0002
#define WM_MOVE                         0x0003
#define WM_SIZE                         0x0005

#define WM_ACTIVATE                     0x0006
 #define     WA_INACTIVE     0
 #define     WA_ACTIVE       1
 #define     WA_CLICKACTIVE  2

 #define WM_SETFOCUS                     0x0007
 #define WM_KILLFOCUS                    0x0008
 #define WM_ENABLE                       0x000A
 #define WM_SETREDRAW                    0x000B
 #define WM_SETTEXT                      0x000C
 #define WM_GETTEXT                      0x000D
 #define WM_GETTEXTLENGTH                0x000E
 #define WM_PAINT                        0x000F
 #define WM_CLOSE                        0x0010
 #ifndef _WIN32_WCE
 #define WM_QUERYENDSESSION              0x0011
 #define WM_QUERYOPEN                    0x0013
 #define WM_ENDSESSION                   0x0016
 #endif
 #define WM_QUIT                         0x0012
 #define WM_ERASEBKGND                   0x0014
 #define WM_SYSCOLORCHANGE               0x0015
 #define WM_SHOWWINDOW                   0x0018
 #define WM_WININICHANGE                 0x001A
 #if(WINVER >= 0x0400)
 #define WM_SETTINGCHANGE                WM_WININICHANGE
 #endif

 #if (NTDDI_VERSION >= NTDDI_WIN10_19H1)
 #endif // NTDDI_VERSION >= NTDDI_WIN10_19H1


 #define WM_DEVMODECHANGE                0x001B
 #define WM_ACTIVATEAPP                  0x001C
 #define WM_FONTCHANGE                   0x001D
 #define WM_TIMECHANGE                   0x001E
 #define WM_CANCELMODE                   0x001F
 #define WM_SETCURSOR                    0x0020
 #define WM_MOUSEACTIVATE                0x0021
 #define WM_CHILDACTIVATE                0x0022
 #define WM_QUEUESYNC                    0x0023

 #define WM_GETMINMAXINFO                0x0024

 #pragma region Desktop Family
 #if WINAPI_FAMILY_PARTITION(WINAPI_PARTITION_DESKTOP)



 #endif
 #pragma endregion

 #define WM_PAINTICON                    0x0026
 #define WM_ICONERASEBKGND               0x0027
 #define WM_NEXTDLGCTL                   0x0028
 #define WM_SPOOLERSTATUS                0x002A
 #define WM_DRAWITEM                     0x002B
 #define WM_MEASUREITEM                  0x002C
 #define WM_DELETEITEM                   0x002D
 #define WM_VKEYTOITEM                   0x002E
 #define WM_CHARTOITEM                   0x002F
 #define WM_SETFONT                      0x0030
 #define WM_GETFONT                      0x0031
 #define WM_SETHOTKEY                    0x0032
 #define WM_GETHOTKEY                    0x0033
 #define WM_QUERYDRAGICON                0x0037
 #define WM_COMPAREITEM                  0x0039
 #if(WINVER >= 0x0500)
 #ifndef _WIN32_WCE
 #define WM_GETOBJECT                    0x003D
 #endif
 #endif
 #define WM_COMPACTING                   0x0041
 #define WM_COMMNOTIFY                   0x0044
 #define WM_WINDOWPOSCHANGING            0x0046
 #define WM_WINDOWPOSCHANGED             0x0047

 #define WM_POWER                        0x0048

 #define PWR_OK              1
 #define PWR_FAIL            (-1)
 #define PWR_SUSPENDREQUEST  1
 #define PWR_SUSPENDRESUME   2
 #define PWR_CRITICALRESUME  3

 #define WM_COPYDATA                     0x004A
 #define WM_CANCELJOURNAL                0x004B


 #pragma region Desktop Family
 #if WINAPI_FAMILY_PARTITION(WINAPI_PARTITION_DESKTOP)


 typedef struct tagCOPYDATASTRUCT {
	 ULONG_PTR dwData;
	 DWORD cbData;
	 _Field_size_bytes_(cbData) PVOID lpData;
 } COPYDATASTRUCT, *PCOPYDATASTRUCT;

 #if(WINVER >= 0x0400)
 typedef struct tagMDINEXTMENU
 {
	 HMENU   hmenuIn;
	 HMENU   hmenuNext;
	 HWND    hwndNext;
 } MDINEXTMENU, * PMDINEXTMENU, FAR * LPMDINEXTMENU;
 #endif

 #endif
 #pragma endregion


 #if(WINVER >= 0x0400)
 #define WM_NOTIFY                       0x004E
 #define WM_INPUTLANGCHANGEREQUEST       0x0050
 #define WM_INPUTLANGCHANGE              0x0051
 #define WM_TCARD                        0x0052
 #define WM_HELP                         0x0053
 #define WM_USERCHANGED                  0x0054
 #define WM_NOTIFYFORMAT                 0x0055

 #define NFR_ANSI                             1
 #define NFR_UNICODE                          2
 #define NF_QUERY                             3
 #define NF_REQUERY                           4

 #define WM_CONTEXTMENU                  0x007B
 #define WM_STYLECHANGING                0x007C
 #define WM_STYLECHANGED                 0x007D
 #define WM_DISPLAYCHANGE                0x007E
 #define WM_GETICON                      0x007F
 #define WM_SETICON                      0x0080
 #endif

 #define WM_NCCREATE                     0x0081
 #define WM_NCDESTROY                    0x0082
 #define WM_NCCALCSIZE                   0x0083
 #define WM_NCHITTEST                    0x0084
 #define WM_NCPAINT                      0x0085
 #define WM_NCACTIVATE                   0x0086
 #define WM_GETDLGCODE                   0x0087
 #ifndef _WIN32_WCE
 #define WM_SYNCPAINT                    0x0088
 #endif
 #define WM_NCMOUSEMOVE                  0x00A0
 #define WM_NCLBUTTONDOWN                0x00A1
 #define WM_NCLBUTTONUP                  0x00A2
 #define WM_NCLBUTTONDBLCLK              0x00A3
 #define WM_NCRBUTTONDOWN                0x00A4
 #define WM_NCRBUTTONUP                  0x00A5
 #define WM_NCRBUTTONDBLCLK              0x00A6
 #define WM_NCMBUTTONDOWN                0x00A7
 #define WM_NCMBUTTONUP                  0x00A8
 #define WM_NCMBUTTONDBLCLK              0x00A9



 #if(_WIN32_WINNT >= 0x0500)
 #define WM_NCXBUTTONDOWN                0x00AB
 #define WM_NCXBUTTONUP                  0x00AC
 #define WM_NCXBUTTONDBLCLK              0x00AD
 #endif


 #if(_WIN32_WINNT >= 0x0501)
 #define WM_INPUT_DEVICE_CHANGE          0x00FE
 #endif

 #if(_WIN32_WINNT >= 0x0501)
 #define WM_INPUT                        0x00FF
 #endif

 #define WM_KEYFIRST                     0x0100
 #define WM_KEYDOWN                      0x0100
 #define WM_KEYUP                        0x0101
 #define WM_CHAR                         0x0102
 #define WM_DEADCHAR                     0x0103
 #define WM_SYSKEYDOWN                   0x0104
 #define WM_SYSKEYUP                     0x0105
 #define WM_SYSCHAR                      0x0106
 #define WM_SYSDEADCHAR                  0x0107
 #if(_WIN32_WINNT >= 0x0501)
 #define WM_UNICHAR                      0x0109
 #define WM_KEYLAST                      0x0109
 #define UNICODE_NOCHAR                  0xFFFF
 #else
 #define WM_KEYLAST                      0x0108
 #endif

 #if(WINVER >= 0x0400)
 #define WM_IME_STARTCOMPOSITION         0x010D
 #define WM_IME_ENDCOMPOSITION           0x010E
 #define WM_IME_COMPOSITION              0x010F
 #define WM_IME_KEYLAST                  0x010F
 #endif

 #define WM_INITDIALOG                   0x0110
 #define WM_COMMAND                      0x0111
 #define WM_SYSCOMMAND                   0x0112
 #define WM_TIMER                        0x0113
 #define WM_HSCROLL                      0x0114
 #define WM_VSCROLL                      0x0115
 #define WM_INITMENU                     0x0116
 #define WM_INITMENUPOPUP                0x0117
 #if(WINVER >= 0x0601)
 #define WM_GESTURE                      0x0119
 #define WM_GESTURENOTIFY                0x011A
 #endif
 #define WM_MENUSELECT                   0x011F
 #define WM_MENUCHAR                     0x0120
 #define WM_ENTERIDLE                    0x0121
 #if(WINVER >= 0x0500)
 #ifndef _WIN32_WCE
 #define WM_MENURBUTTONUP                0x0122
 #define WM_MENUDRAG                     0x0123
 #define WM_MENUGETOBJECT                0x0124
 #define WM_UNINITMENUPOPUP              0x0125
 #define WM_MENUCOMMAND                  0x0126

 #ifndef _WIN32_WCE
 #if(_WIN32_WINNT >= 0x0500)
 #define WM_CHANGEUISTATE                0x0127
 #define WM_UPDATEUISTATE                0x0128
 #define WM_QUERYUISTATE                 0x0129


 #define UIS_SET                         1
 #define UIS_CLEAR                       2
 #define UIS_INITIALIZE                  3


 #define UISF_HIDEFOCUS                  0x1
 #define UISF_HIDEACCEL                  0x2
 #if(_WIN32_WINNT >= 0x0501)
 #define UISF_ACTIVE                     0x4
 #endif
 #endif
 #endif

 #endif
 #endif

 #define WM_CTLCOLORMSGBOX               0x0132
 #define WM_CTLCOLOREDIT                 0x0133
 #define WM_CTLCOLORLISTBOX              0x0134
 #define WM_CTLCOLORBTN                  0x0135
 #define WM_CTLCOLORDLG                  0x0136
 #define WM_CTLCOLORSCROLLBAR            0x0137
 #define WM_CTLCOLORSTATIC               0x0138
 #define MN_GETHMENU                     0x01E1

 #define WM_MOUSEFIRST                   0x0200
 #define WM_MOUSEMOVE                    0x0200
 #define WM_LBUTTONDOWN                  0x0201
 #define WM_LBUTTONUP                    0x0202
 #define WM_LBUTTONDBLCLK                0x0203
 #define WM_RBUTTONDOWN                  0x0204
 #define WM_RBUTTONUP                    0x0205
 #define WM_RBUTTONDBLCLK                0x0206
 #define WM_MBUTTONDOWN                  0x0207
 #define WM_MBUTTONUP                    0x0208
 #define WM_MBUTTONDBLCLK                0x0209
 #if (_WIN32_WINNT >= 0x0400) || (_WIN32_WINDOWS > 0x0400)
 #define WM_MOUSEWHEEL                   0x020A
 #endif
 #if (_WIN32_WINNT >= 0x0500)
 #define WM_XBUTTONDOWN                  0x020B
 #define WM_XBUTTONUP                    0x020C
 #define WM_XBUTTONDBLCLK                0x020D
 #endif
 #if (_WIN32_WINNT >= 0x0600)
 #define WM_MOUSEHWHEEL                  0x020E
 #endif

 #if (_WIN32_WINNT >= 0x0600)
 #define WM_MOUSELAST                    0x020E
 #elif (_WIN32_WINNT >= 0x0500)
 #define WM_MOUSELAST                    0x020D
 #elif (_WIN32_WINNT >= 0x0400) || (_WIN32_WINDOWS > 0x0400)
 #define WM_MOUSELAST                    0x020A
 #else
 #define WM_MOUSELAST                    0x0209
 #endif


 #if(_WIN32_WINNT >= 0x0400)

 #define WHEEL_DELTA                     120
 #define GET_WHEEL_DELTA_WPARAM(wParam)  ((short)HIWORD(wParam))


 #define WHEEL_PAGESCROLL                (UINT_MAX)
 #endif

 #if(_WIN32_WINNT >= 0x0500)
 #define GET_KEYSTATE_WPARAM(wParam)     (LOWORD(wParam))
 #define GET_NCHITTEST_WPARAM(wParam)    ((short)LOWORD(wParam))
 #define GET_XBUTTON_WPARAM(wParam)      (HIWORD(wParam))


 #define XBUTTON1      0x0001
 #define XBUTTON2      0x0002

 #endif

 #define WM_PARENTNOTIFY                 0x0210
 #define WM_ENTERMENULOOP                0x0211
 #define WM_EXITMENULOOP                 0x0212

 #if(WINVER >= 0x0400)
 #define WM_NEXTMENU                     0x0213
 #define WM_SIZING                       0x0214
 #define WM_CAPTURECHANGED               0x0215
 #define WM_MOVING                       0x0216
 #endif

 #if(WINVER >= 0x0400)


 #define WM_POWERBROADCAST               0x0218

 #ifndef _WIN32_WCE
 #define PBT_APMQUERYSUSPEND             0x0000
 #define PBT_APMQUERYSTANDBY             0x0001

 #define PBT_APMQUERYSUSPENDFAILED       0x0002
 #define PBT_APMQUERYSTANDBYFAILED       0x0003

 #define PBT_APMSUSPEND                  0x0004
 #define PBT_APMSTANDBY                  0x0005

 #define PBT_APMRESUMECRITICAL           0x0006
 #define PBT_APMRESUMESUSPEND            0x0007
 #define PBT_APMRESUMESTANDBY            0x0008

 #define PBTF_APMRESUMEFROMFAILURE       0x00000001

 #define PBT_APMBATTERYLOW               0x0009
 #define PBT_APMPOWERSTATUSCHANGE        0x000A

 #define PBT_APMOEMEVENT                 0x000B


 #define PBT_APMRESUMEAUTOMATIC          0x0012
 #if (_WIN32_WINNT >= 0x0502)
 #ifndef PBT_POWERSETTINGCHANGE
 #define PBT_POWERSETTINGCHANGE          0x8013

 #pragma region Desktop Family
 #if WINAPI_FAMILY_PARTITION(WINAPI_PARTITION_DESKTOP)

 typedef struct {
	 GUID PowerSetting;
	 DWORD DataLength;
	 UCHAR Data[1];
 } POWERBROADCAST_SETTING, *PPOWERBROADCAST_SETTING;


 #endif
 #pragma endregion

 #endif // PBT_POWERSETTINGCHANGE

 #endif // (_WIN32_WINNT >= 0x0502)
 #endif

 #endif

 #if(WINVER >= 0x0400)
 #define WM_DEVICECHANGE                 0x0219
 #endif

 #define WM_MDICREATE                    0x0220
 #define WM_MDIDESTROY                   0x0221
 #define WM_MDIACTIVATE                  0x0222
 #define WM_MDIRESTORE                   0x0223
 #define WM_MDINEXT                      0x0224
 #define WM_MDIMAXIMIZE                  0x0225
 #define WM_MDITILE                      0x0226
 #define WM_MDICASCADE                   0x0227
 #define WM_MDIICONARRANGE               0x0228
 #define WM_MDIGETACTIVE                 0x0229


 #define WM_MDISETMENU                   0x0230
 #define WM_ENTERSIZEMOVE                0x0231
 #define WM_EXITSIZEMOVE                 0x0232
 #define WM_DROPFILES                    0x0233
 #define WM_MDIREFRESHMENU               0x0234

 #if(WINVER >= 0x0602)
 #define WM_POINTERDEVICECHANGE          0x238
 #define WM_POINTERDEVICEINRANGE         0x239
 #define WM_POINTERDEVICEOUTOFRANGE      0x23A
 #endif


 #if(WINVER >= 0x0601)
 #define WM_TOUCH                        0x0240
 #endif

 #if(WINVER >= 0x0602)
 #define WM_NCPOINTERUPDATE              0x0241
 #define WM_NCPOINTERDOWN                0x0242
 #define WM_NCPOINTERUP                  0x0243
 #define WM_POINTERUPDATE                0x0245
 #define WM_POINTERDOWN                  0x0246
 #define WM_POINTERUP                    0x0247
 #define WM_POINTERENTER                 0x0249
 #define WM_POINTERLEAVE                 0x024A
 #define WM_POINTERACTIVATE              0x024B
 #define WM_POINTERCAPTURECHANGED        0x024C
 #define WM_TOUCHHITTESTING              0x024D
 #define WM_POINTERWHEEL                 0x024E
 #define WM_POINTERHWHEEL                0x024F
 #define DM_POINTERHITTEST               0x0250
 #define WM_POINTERROUTEDTO              0x0251
 #define WM_POINTERROUTEDAWAY            0x0252
 #define WM_POINTERROUTEDRELEASED        0x0253
 #endif


 #if(WINVER >= 0x0400)
 #define WM_IME_SETCONTEXT               0x0281
 #define WM_IME_NOTIFY                   0x0282
 #define WM_IME_CONTROL                  0x0283
 #define WM_IME_COMPOSITIONFULL          0x0284
 #define WM_IME_SELECT                   0x0285
 #define WM_IME_CHAR                     0x0286
 #endif
 #if(WINVER >= 0x0500)
 #define WM_IME_REQUEST                  0x0288
 #endif
 #if(WINVER >= 0x0400)
 #define WM_IME_KEYDOWN                  0x0290
 #define WM_IME_KEYUP                    0x0291
 #endif

 #if((_WIN32_WINNT >= 0x0400) || (WINVER >= 0x0500))
 #define WM_MOUSEHOVER                   0x02A1
 #define WM_MOUSELEAVE                   0x02A3
 #endif
 #if(WINVER >= 0x0500)
 #define WM_NCMOUSEHOVER                 0x02A0
 #define WM_NCMOUSELEAVE                 0x02A2
 #endif

 #if(_WIN32_WINNT >= 0x0501)
 #define WM_WTSSESSION_CHANGE            0x02B1

 #define WM_TABLET_FIRST                 0x02c0
 #define WM_TABLET_LAST                  0x02df
 #endif

 #if(WINVER >= 0x0601)
 #define WM_DPICHANGED                   0x02E0
 #endif
 #if(WINVER >= 0x0605)
 #define WM_DPICHANGED_BEFOREPARENT      0x02E2
 #define WM_DPICHANGED_AFTERPARENT       0x02E3
 #define WM_GETDPISCALEDSIZE             0x02E4
 #endif

 #define WM_CUT                          0x0300
 #define WM_COPY                         0x0301
 #define WM_PASTE                        0x0302
 #define WM_CLEAR                        0x0303
 #define WM_UNDO                         0x0304
 #define WM_RENDERFORMAT                 0x0305
 #define WM_RENDERALLFORMATS             0x0306
 #define WM_DESTROYCLIPBOARD             0x0307
 #define WM_DRAWCLIPBOARD                0x0308
 #define WM_PAINTCLIPBOARD               0x0309
 #define WM_VSCROLLCLIPBOARD             0x030A
 #define WM_SIZECLIPBOARD                0x030B
 #define WM_ASKCBFORMATNAME              0x030C
 #define WM_CHANGECBCHAIN                0x030D
 #define WM_HSCROLLCLIPBOARD             0x030E
 #define WM_QUERYNEWPALETTE              0x030F
 #define WM_PALETTEISCHANGING            0x0310
 #define WM_PALETTECHANGED               0x0311
 #define WM_HOTKEY                       0x0312

 #if(WINVER >= 0x0400)
 #define WM_PRINT                        0x0317
 #define WM_PRINTCLIENT                  0x0318
 #endif

 #if(_WIN32_WINNT >= 0x0500)
 #define WM_APPCOMMAND                   0x0319
 #endif

 #if(_WIN32_WINNT >= 0x0501)
 #define WM_THEMECHANGED                 0x031A
 #endif


 #if(_WIN32_WINNT >= 0x0501)
 #define WM_CLIPBOARDUPDATE              0x031D
 #endif

 #if(_WIN32_WINNT >= 0x0600)
 #define WM_DWMCOMPOSITIONCHANGED        0x031E
 #define WM_DWMNCRENDERINGCHANGED        0x031F
 #define WM_DWMCOLORIZATIONCOLORCHANGED  0x0320
 #define WM_DWMWINDOWMAXIMIZEDCHANGE     0x0321
 #endif

 #if(_WIN32_WINNT >= 0x0601)
 #define WM_DWMSENDICONICTHUMBNAIL           0x0323
 #define WM_DWMSENDICONICLIVEPREVIEWBITMAP   0x0326
 #endif


 #if(WINVER >= 0x0600)
 #define WM_GETTITLEBARINFOEX            0x033F
 #endif

 #if(WINVER >= 0x0400)
 #endif


 #if(WINVER >= 0x0400)
 #define WM_HANDHELDFIRST                0x0358
 #define WM_HANDHELDLAST                 0x035F

 #define WM_AFXFIRST                     0x0360
 #define WM_AFXLAST                      0x037F
 #endif

 #define WM_PENWINFIRST                  0x0380
 #define WM_PENWINLAST                   0x038F
*/
var messageNames = map[uint32]string{
	// --- Window ---
	0x0000: "WM_NULL",
	0x0001: "WM_CREATE",
	0x0002: "WM_DESTROY",
	0x0003: "WM_MOVE",
	0x0005: "WM_SIZE",
	0x0006: "WM_ACTIVATE",
	0x0007: "WM_SETFOCUS",
	0x0008: "WM_KILLFOCUS",
	0x000A: "WM_ENABLE",
	0x000B: "WM_SETREDRAW",
	0x000C: "WM_SETTEXT",
	0x000D: "WM_GETTEXT",
	0x000E: "WM_GETTEXTLENGTH",
	0x000F: "WM_PAINT",
	0x0010: "WM_CLOSE",
	0x0011: "WM_QUERYENDSESSION",
	0x0012: "WM_QUIT",
	0x0014: "WM_ERASEBKGND",
	0x0015: "WM_SYSCOLORCHANGE",
	0x0018: "WM_SHOWWINDOW",
	0x001A: "WM_WININICHANGE",
	0x001B: "WM_DEVMODECHANGE",
	0x001C: "WM_ACTIVATEAPP",
	0x001F: "WM_CANCELMODE",
	0x0020: "WM_SETCURSOR",
	0x0021: "WM_MOUSEACTIVATE",
	0x0024: "WM_GETMINMAXINFO",
	0x0026: "WM_PAINTICON",
	0x0027: "WM_ICONERASEBKGND",
	0x0028: "WM_NEXTDLGCTL",
	0x002A: "WM_SPOOLERSTATUS",
	0x002B: "WM_DRAWITEM",
	0x002C: "WM_MEASUREITEM",
	0x002D: "WM_DELETEITEM",
	0x002E: "WM_VKEYTOITEM",
	0x002F: "WM_CHARTOITEM",
	0x0030: "WM_SETFONT",
	0x0031: "WM_GETFONT",
	0x0032: "WM_SETHOTKEY",
	0x0033: "WM_GETHOTKEY",
	0x0037: "WM_QUERYDRAGICON",
	0x0039: "WM_COMPAREITEM",
	0x003D: "WM_GETOBJECT",
	0x0041: "WM_COMPACTING",
	0x0044: "WM_COMMNOTIFY",
	0x0046: "WM_WINDOWPOSCHANGING",
	0x0047: "WM_WINDOWPOSCHANGED",
	0x0048: "WM_POWER",
	0x004A: "WM_COPYDATA",
	0x004B: "WM_CANCELJOURNAL",
	0x007B: "WM_CONTEXTMENU",
	0x007C: "WM_STYLECHANGING",
	0x007D: "WM_STYLECHANGED",
	0x007E: "WM_DISPLAYCHANGE",
	0x007F: "WM_GETICON",
	0x0080: "WM_SETICON",
	0x0081: "WM_NCCREATE",
	0x0082: "WM_NCDESTROY",
	0x0083: "WM_NCCALCSIZE",
	0x0084: "WM_NCHITTEST",
	0x0085: "WM_NCPAINT",
	0x0086: "WM_NCACTIVATE",
	0x0087: "WM_GETDLGCODE",
	0x0088: "WM_SYNCPAINT",
	0x00A0: "WM_NCMOUSEMOVE",
	0x00A1: "WM_NCLBUTTONDOWN",
	0x00A2: "WM_NCLBUTTONUP",
	0x00A3: "WM_NCLBUTTONDBLCLK",
	0x00A4: "WM_NCRBUTTONDOWN",
	0x00A5: "WM_NCRBUTTONUP",
	0x00A6: "WM_NCRBUTTONDBLCLK",
	0x00A7: "WM_NCMBUTTONDOWN",
	0x00A8: "WM_NCMBUTTONUP",
	0x00A9: "WM_NCMBUTTONDBLCLK",
	0x00AB: "WM_NCXBUTTONDOWN",
	0x00AC: "WM_NCXBUTTONUP",
	0x00AD: "WM_NCXBUTTONDBLCLK",
	0x00FE: "WM_INPUT_DEVICE_CHANGE",
	0x00FF: "WM_INPUT",
	0x0100: "WM_KEYDOWN",
	0x0101: "WM_KEYUP",
	0x0102: "WM_CHAR",
	0x0103: "WM_DEADCHAR",
	0x0104: "WM_SYSKEYDOWN",
	0x0105: "WM_SYSKEYUP",
	0x0106: "WM_SYSCHAR",
	0x0107: "WM_SYSDEADCHAR",
	0x0109: "WM_UNICHAR",
	0x010D: "WM_IME_STARTCOMPOSITION",
	0x010E: "WM_IME_ENDCOMPOSITION",
	0x010F: "WM_IME_COMPOSITION",
	0x0110: "WM_INITDIALOG",
	0x0111: "WM_COMMAND",
	0x0112: "WM_SYSCOMMAND",
	0x0113: "WM_TIMER",
	0x0114: "WM_HSCROLL",
	0x0115: "WM_VSCROLL",
	0x0116: "WM_INITMENU",
	0x0117: "WM_INITMENUPOPUP",
	0x0119: "WM_GESTURE",
	0x011A: "WM_GESTURENOTIFY",
	0x011F: "WM_MENUSELECT",
	0x0120: "WM_MENUCHAR",
	0x0121: "WM_ENTERIDLE",
	0x0122: "WM_MENURBUTTONUP",
	0x0123: "WM_MENUDRAG",
	0x0124: "WM_MENUGETOBJECT",
	0x0125: "WM_UNINITMENUPOPUP",
	0x0126: "WM_MENUCOMMAND",
	0x0127: "WM_CHANGEUISTATE",
	0x0128: "WM_UPDATEUISTATE",
	0x0129: "WM_QUERYUISTATE",
	0x012B: "WM_CTLCOLORMSGBOX",
	0x012C: "WM_CTLCOLOREDIT",
	0x012D: "WM_CTLCOLORLISTBOX",
	0x012E: "WM_CTLCOLORBTN",
	0x012F: "WM_CTLCOLORDLG",
	0x0130: "WM_CTLCOLORSCROLLBAR",
	0x0131: "WM_CTLCOLORSTATIC",
	0x01E1: "MN_GETHMENU",
	0x0200: "WM_MOUSEMOVE",
	0x0201: "WM_LBUTTONDOWN",
	0x0202: "WM_LBUTTONUP",
	0x0203: "WM_LBUTTONDBLCLK",
	0x0204: "WM_RBUTTONDOWN",
	0x0205: "WM_RBUTTONUP",
	0x0206: "WM_RBUTTONDBLCLK",
	0x0207: "WM_MBUTTONDOWN",
	0x0208: "WM_MBUTTONUP",
	0x0209: "WM_MBUTTONDBLCLK",
	0x020A: "WM_MOUSEWHEEL",
	0x020B: "WM_XBUTTONDOWN",
	0x020C: "WM_XBUTTONUP",
	0x020D: "WM_XBUTTONDBLCLK",
	0x020E: "WM_MOUSEHWHEEL",
	0x0210: "WM_PARENTNOTIFY",
	0x0211: "WM_ENTERMENULOOP",
	0x0212: "WM_EXITMENULOOP",
	0x0213: "WM_NEXTMENU",
	0x0214: "WM_SIZING",
	0x0215: "WM_CAPTURECHANGED",
	0x0216: "WM_MOVING",
	0x0218: "WM_POWERBROADCAST",
	0x0219: "WM_DEVICECHANGE",
	0x0220: "WM_MDICREATE",
	0x0221: "WM_MDIDESTROY",
	0x0222: "WM_MDIACTIVATE",
	0x0223: "WM_MDIRESTORE",
	0x0224: "WM_MDINEXT",
	0x0225: "WM_MDIMAXIMIZE",
	0x0226: "WM_MDITILE",
	0x0227: "WM_MDICASCADE",
	0x0228: "WM_MDIICONARRANGE",
	0x0229: "WM_MDIGETACTIVE",
	0x0230: "WM_MDISETMENU",
	0x0231: "WM_ENTERSIZEMOVE",
	0x0232: "WM_EXITSIZEMOVE",
	0x0233: "WM_DROPFILES",
	0x0234: "WM_MDIREFRESHMENU",
	0x0240: "WM_TOUCH",
	0x0241: "WM_NCPOINTERUPDATE",
	0x0242: "WM_NCPOINTERDOWN",
	0x0243: "WM_NCPOINTERUP",
	0x0245: "WM_POINTERUPDATE",
	0x0246: "WM_POINTERDOWN",
	0x0247: "WM_POINTERUP",
	0x0249: "WM_POINTERENTER",
	0x024A: "WM_POINTERLEAVE",
	0x024B: "WM_POINTERACTIVATE",
	0x024C: "WM_POINTERCAPTURECHANGED",
	0x024D: "WM_TOUCHHITTESTING",
	0x024E: "WM_POINTERWHEEL",
	0x024F: "WM_POINTERHWHEEL",
	0x0250: "DM_POINTERHITTEST",
	0x0251: "WM_POINTERROUTEDTO",
	0x0252: "WM_POINTERROUTEDAWAY",
	0x0253: "WM_POINTERROUTEDRELEASED",
	0x0281: "WM_IME_SETCONTEXT",
	0x0282: "WM_IME_NOTIFY",
	0x0283: "WM_IME_CONTROL",
	0x0284: "WM_IME_COMPOSITIONFULL",
	0x0285: "WM_IME_SELECT",
	0x0286: "WM_IME_CHAR",
	0x0288: "WM_IME_REQUEST",
	0x0290: "WM_IME_KEYDOWN",
	0x0291: "WM_IME_KEYUP",
	0x02A0: "WM_NCMOUSEHOVER",
	0x02A1: "WM_MOUSEHOVER",
	0x02A2: "WM_NCMOUSELEAVE",
	0x02A3: "WM_MOUSELEAVE",
	0x02B1: "WM_WTSSESSION_CHANGE",
	0x02C0: "WM_TABLET_FIRST",
	0x02DF: "WM_TABLET_LAST",
	0x02E0: "WM_DPICHANGED",
	0x02E2: "WM_DPICHANGED_BEFOREPARENT",
	0x02E3: "WM_DPICHANGED_AFTERPARENT",
	0x02E4: "WM_GETDPISCALEDSIZE",
	0x0300: "WM_CUT",
	0x0301: "WM_COPY",
	0x0302: "WM_PASTE",
	0x0303: "WM_CLEAR",
	0x0304: "WM_UNDO",
	0x0305: "WM_RENDERFORMAT",
	0x0306: "WM_RENDERALLFORMATS",
	0x0307: "WM_DESTROYCLIPBOARD",
	0x0308: "WM_DRAWCLIPBOARD",
	0x0309: "WM_PAINTCLIPBOARD",
	0x030A: "WM_VSCROLLCLIPBOARD",
	0x030B: "WM_SIZECLIPBOARD",
	0x030C: "WM_ASKCBFORMATNAME",
	0x030D: "WM_CHANGECBCHAIN",
	0x030E: "WM_HSCROLLCLIPBOARD",
	0x030F: "WM_QUERYNEWPALETTE",
	0x0310: "WM_PALETTEISCHANGING",
	0x0311: "WM_PALETTECHANGED",
	0x0312: "WM_HOTKEY",
	0x0317: "WM_PRINT",
	0x0318: "WM_PRINTCLIENT",
	0x0319: "WM_APPCOMMAND",
	0x031A: "WM_THEMECHANGED",
	0x031D: "WM_CLIPBOARDUPDATE",
	0x031E: "WM_DWMCOMPOSITIONCHANGED",
	0x031F: "WM_DWMNCRENDERINGCHANGED",
	0x0320: "WM_DWMCOLORIZATIONCOLORCHANGED",
	0x0321: "WM_DWMWINDOWMAXIMIZEDCHANGE",
	0x0323: "WM_DWMSENDICONICTHUMBNAIL",
	0x0326: "WM_DWMSENDICONICLIVEPREVIEWBITMAP",
	0x033F: "WM_GETTITLEBARINFOEX",
	0x0358: "WM_HANDHELDFIRST",
	0x035F: "WM_HANDHELDLAST",
	0x0360: "WM_AFXFIRST",
	0x037F: "WM_AFXLAST",
	0x0380: "WM_PENWINFIRST",
	0x038F: "WM_PENWINLAST",
	0x0400: "WM_USER",
	0x8000: "WM_APP",
}

// MessageName returns a string like "WM_KEYDOWN" or "UNKNOWN (0x0100)"
func MessageName(msg uint32) string {
	if name, ok := messageNames[msg]; ok {
		return name
	}
	return "UNKNOWN (0x" + itox(msg) + ")"
}

func itox(v uint32) string {
	const hex = "0123456789ABCDEF"
	b := [10]byte{'0', 'x'}
	for i := 7; i >= 0; i-- {
		b[9-i] = hex[(v>>(i*4))&0xF]
	}
	return string(b[:])
}

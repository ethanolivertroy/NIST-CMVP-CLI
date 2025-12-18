Output demo.gif

Set FontSize 16
Set Width 1200
Set Height 800
Set Padding 20
Set FontFamily "JetBrains Mono"
Set Theme { "name": "CMVP-TUI", "background": "#1E1F29", "foreground": "#FFFFFF", "black": "#282A36", "red": "#FF5F56", "green": "#04B575", "yellow": "#FFCC00", "blue": "#45B7D1", "magenta": "#7D56F4", "cyan": "#4ECDC4", "white": "#F8F8F2", "brightBlack": "#626262", "brightRed": "#FF6B6B", "brightGreen": "#58D68D", "brightYellow": "#F7DC6F", "brightBlue": "#85C1E9", "brightMagenta": "#BB8FCE", "brightCyan": "#98D8C8", "brightWhite": "#FFFFFF", "cursor": "#7D56F4", "selection": "#44475A" }
Set TypingSpeed 50ms

# ══════════════════════════════════════════════════════════════
# 1. LAUNCH & LOADING
# ══════════════════════════════════════════════════════════════
Hide
Type "./cmvp-tui"
Enter
Sleep 5s
Show

Sleep 2s

# ══════════════════════════════════════════════════════════════
# 2. LIST NAVIGATION - Browse modules with status badges
# ══════════════════════════════════════════════════════════════

Type "j"
Sleep 400ms
Type "j"
Sleep 400ms
Type "j"
Sleep 400ms
Type "j"
Sleep 400ms
Type "j"
Sleep 400ms
Type "j"
Sleep 1.5s

# ══════════════════════════════════════════════════════════════
# 3. SEARCH DEMONSTRATIONS - Popular vendors
# ══════════════════════════════════════════════════════════════

# Search for OpenSSL modules
Type "/"
Sleep 300ms
Type@100ms "openssl"
Sleep 2.5s
Escape
Sleep 500ms

# Search for Amazon/AWS modules
Type "/"
Sleep 300ms
Type@100ms "amazon"
Sleep 2.5s
Escape
Sleep 500ms

# Search for BoringCrypto (Google)
Type "/"
Sleep 300ms
Type@100ms "boringcrypto"
Sleep 2.5s
Escape
Sleep 500ms

# Search for wolfCrypt
Type "/"
Sleep 300ms
Type@100ms "wolfcrypt"
Sleep 2.5s
Escape
Sleep 500ms

# Search by certificate number (show flexibility)
Type "/"
Sleep 300ms
Type@100ms "5100"
Sleep 2.5s
Escape
Sleep 500ms

# ══════════════════════════════════════════════════════════════
# 4. DETAIL VIEW - Ubuntu 22.04 OpenSSL Module
#    Shows: status badge, security level, caveat, algorithms
# ══════════════════════════════════════════════════════════════

Type "/"
Sleep 300ms
Type@100ms "Ubuntu 22.04 OpenSSL"
Sleep 1.5s
Enter
Sleep 500ms
Enter
Sleep 3s

# Toggle to detailed algorithm view
Type "d"
Sleep 2s

# Scroll through algorithms
Type "j"
Sleep 400ms
Type "j"
Sleep 400ms
Type "j"
Sleep 400ms
Type "j"
Sleep 1.5s

# Toggle back to category view
Type "d"
Sleep 1.5s

# Return to list
Escape
Sleep 500ms

# CRITICAL: Clear filter before next search
Escape
Sleep 500ms

# ══════════════════════════════════════════════════════════════
# 5. DETAIL VIEW - wolfCrypt (popular module)
# ══════════════════════════════════════════════════════════════

Type "/"
Sleep 300ms
Type@100ms "wolfCrypt"
Sleep 1.5s
Enter
Sleep 500ms
Enter
Sleep 3s

# Return to list
Escape
Sleep 500ms

# ══════════════════════════════════════════════════════════════
# 6. FINAL - Clear and exit (smooth loop)
# ══════════════════════════════════════════════════════════════

Escape
Sleep 500ms
Hide
Type "q"
Sleep 500ms

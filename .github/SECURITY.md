# Security Policy

## Supported Versions

| Version | Supported          |
| ------- | ------------------ |
| latest  | :white_check_mark: |

## Reporting a Vulnerability

If you discover a security vulnerability in cmvp-tui, please report it responsibly:

1. **Do not** open a public GitHub issue for security vulnerabilities
2. Email the maintainer directly or use [GitHub's private vulnerability reporting](https://github.com/ethanolivertroy/cmvp-tui/security/advisories/new)
3. Include:
   - Description of the vulnerability
   - Steps to reproduce
   - Potential impact
   - Suggested fix (if any)

## Response Timeline

- **Acknowledgment**: Within 48 hours
- **Initial assessment**: Within 7 days
- **Fix timeline**: Depends on severity, typically within 30 days for critical issues

## Security Measures

This project implements the following security practices:

- Dependency scanning via OSV-Scanner and govulncheck
- OpenSSF Scorecard monitoring
- SHA-pinned GitHub Actions
- Minimal permissions in CI workflows

## Scope

This security policy applies to the cmvp-tui CLI tool. The data displayed comes from the public NIST CMVP database and is not modified by this tool.

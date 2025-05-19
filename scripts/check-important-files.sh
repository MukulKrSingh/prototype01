#!/bin/bash

# Script to check for the existence of important project files
# Usage: ./scripts/check-important-files.sh

# List of critical files that should never be deleted
CRITICAL_FILES=(
  "COPILOT_INSTRUCTIONS.md"
  "README.md"
  "SETUP_SUMMARY.md"
  "PROJECT_STRUCTURE.md"
)

echo "Checking for critical project files..."
MISSING=0

for file in "${CRITICAL_FILES[@]}"; do
  if [ ! -f "$file" ]; then
    echo "⚠️  WARNING: Critical file '$file' is missing!"
    MISSING=1
  else
    echo "✅ Found critical file: $file"
  fi
done

if [ $MISSING -eq 1 ]; then
  echo "❌ Some critical files are missing. Please restore them from version control or backups."
  echo "These files contain essential project information and development guidelines."
  exit 1
else
  echo "✅ All critical project files are present."
  exit 0
fi

#!/usr/bin/env python3
"""Build the frontend into the backend's embedded static dir, then compile the backend."""

import os
import shutil
import subprocess
import sys
from pathlib import Path

ROOT = Path(__file__).resolve().parent


def run(argv: list[str], cwd: Path) -> None:
    exe = shutil.which(argv[0])
    if exe is None:
        sys.exit(f"error: {argv[0]} not found in PATH")

    # Batch shims such as pnpm.cmd cannot be executed directly on Windows.
    command = [exe, *argv[1:]]
    if exe.lower().endswith((".cmd", ".bat")):
        command = [os.environ.get("COMSPEC", "cmd.exe"), "/c", *command]

    print(f"+ {' '.join(argv)}", flush=True)
    subprocess.run(command, cwd=cwd, check=True)


def main() -> None:
    # The frontend bundles into backend/route/static_route/www_root, which the backend
    # embeds at compile time, so the backend must be built after the frontend.
    run(["pnpm", "build"], ROOT / "frontend")
    run(
        ["go", "build", "-tags", "nomsgpack", "-trimpath", "-ldflags", "-s -w"],
        ROOT / "backend",
    )


if __name__ == "__main__":
    try:
        main()
    except subprocess.CalledProcessError as exc:
        sys.exit(exc.returncode)
    except KeyboardInterrupt:
        sys.exit(130)

package ipc

var PathFormats = map[string][]string{
	"windows": {
		`\\?\pipe\discord-ipc-`,
	},
	"linux": {
		`${XDG_RUNTIME_DIR}/discord-ipc-`,
		`${TMPDIR}/discord-ipc-`,
		`${TMP}/discord-ipc-`,
		`${TEMP}/discord-ipc-`,
		`/tmp/discord-ipc-`,
	},
	"darwin": {
		`${XDG_RUNTIME_DIR}/discord-ipc-`,
		`${TMPDIR}/discord-ipc-`,
		`${TMP}/discord-ipc-`,
		`${TEMP}/discord-ipc-`,
		`/tmp/discord-ipc-`,
	},
}

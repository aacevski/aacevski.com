import { readFileSync } from "node:fs";
import { join } from "node:path";

export function loadAvatarImage() {
	const image = readFileSync(
		join(process.cwd(), "src", "assets", "andrej.png"),
	);

	return `data:image/png;base64,${image.toString("base64")}`;
}

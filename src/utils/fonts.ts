import { readFileSync } from "node:fs";
import { join } from "node:path";

export function loadFonts(): {
	name: string;
	data: Buffer;
	weight: 100 | 200 | 300 | 400 | 500 | 600 | 700 | 800 | 900;
	style: "normal" | "italic";
}[] {
	const interRegular = readFileSync(
		join(process.cwd(), "public", "fonts", "Inter-Regular.ttf"),
	);
	const jetBrainsMono = readFileSync(
		join(process.cwd(), "public", "fonts", "JetBrainsMono-Bold.ttf"),
	);

	return [
		{
			name: "Inter",
			data: interRegular,
			weight: 100,
			style: "normal",
		},
		{
			name: "JetBrains Mono",
			data: jetBrainsMono,
			weight: 700,
			style: "normal",
		},
	];
}

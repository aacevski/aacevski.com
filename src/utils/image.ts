import { readFileSync } from 'fs';
import { join } from 'path';

export function loadAvatarImage() {
      const image = readFileSync(
        join(process.cwd(), 'src', 'assets', 'andrej.png')
      );

      return `data:image/png;base64,${image.toString('base64')}`;
}
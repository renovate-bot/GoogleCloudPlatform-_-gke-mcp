import fs from 'fs';
import path from 'path';
import { fileURLToPath } from 'url';
import { execSync } from 'child_process';

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);
const uiDir = path.join(__dirname, '..');
const appsDir = path.join(uiDir, 'apps');

const apps: string[] = fs.readdirSync(appsDir).filter((file: string) => {
  if (file === '.' || file === '..') {
    return false;
  }
  try {
    return fs.statSync(path.join(appsDir, file)).isDirectory();
  } catch {
    return false;
  }
});

console.log(`Found apps: ${apps.join(', ')}`);

for (const app of apps) {
  console.log(`\n--- Building app: ${app} ---\n`);
  execSync('npx vite build', {
    stdio: 'inherit',
    cwd: uiDir,
    env: {
      ...process.env,
      VITE_APP_NAME: app,
    },
  });
}

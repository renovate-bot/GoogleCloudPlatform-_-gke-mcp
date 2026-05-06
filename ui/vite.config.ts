import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';
import { viteSingleFile } from 'vite-plugin-singlefile';
import { resolve, dirname } from 'path';
import { fileURLToPath } from 'url';

const __dirname = dirname(fileURLToPath(import.meta.url));
const fastDeepAssignModule = '/node_modules/@mui/utils/fastDeepAssign/fastDeepAssign.';

const appName = process.env.VITE_APP_NAME;

if (!appName) {
  throw new Error('VITE_APP_NAME is not defined');
}

const isDevelopment = process.env.NODE_ENV === 'development';

export default defineConfig({
  plugins: [
    {
      name: 'patch-mui-fast-deep-assign',
      transform(code, id) {
        if (!id.replaceAll('\\', '/').includes(fastDeepAssignModule)) {
          return null;
        }

        // Guard recursive assignment against inherited prototype properties.
        return code.replace(
          /if\s*\(\s*key\s+in\s+target\s*\)\s*\{/,
          'if (Object.prototype.hasOwnProperty.call(target, key)) {'
        );
      },
    },
    react(),
    viteSingleFile(),
  ],
  resolve: {
    alias: {
      '@gke-mcp/ui/shared': resolve(__dirname, 'shared'),
    },
  },
  build: {
    outDir: 'dist',
    emptyOutDir: false,
    sourcemap: isDevelopment ? 'inline' : undefined,
    cssMinify: !isDevelopment,
    minify: !isDevelopment,

    rollupOptions: {
      input: {
        [appName]: resolve(__dirname, `apps/${appName}/index.html`),
      },
    },
  },
});

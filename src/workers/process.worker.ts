/// <reference lib="webworker" />

import { setupWasm } from 'src/wasm/wasm_exec';

let initialized = false;

addEventListener('message', async ({ data: { command, payload } }) => {
    if (!initialized) {
        if (!initialized) {
            await loadWasm();
            initialized = true;
        }
        postMessage({ status: 'initialized' });
    }

    try {
        if (command === 'processImage') {
            const { byteArray, threshold, size } = payload;

            const result = (self as any).processImage(byteArray, threshold, size);
            postMessage({ status: 'completed', result });
        } else if (command === 'processText') {
            const { text, outline, threshold, size } = payload;

            const result = (self as any).processText(text, outline, threshold, size);
            postMessage({ status: 'completed', result });
        }
    } catch (error: any) {
        postMessage({ status: 'error', error: error.message });
    }
});

async function loadWasm(): Promise<void> {
    setupWasm();
    const go = new Go();
    const result = await WebAssembly.instantiateStreaming(fetch('assets/main.wasm'), go.importObject);
    go.run(result.instance);
}

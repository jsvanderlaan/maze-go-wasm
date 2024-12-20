import { Injectable } from '@angular/core';
import { Settings } from 'src/app/settings.component.js';
import { setupWasm } from '../wasm/wasm_exec.js';

@Injectable({ providedIn: 'root' })
export class WasmService {
    async loadAndRunGoWasm(): Promise<void> {
        setupWasm();
        const go = new Go();
        const result = await WebAssembly.instantiateStreaming(fetch('assets/main.wasm'), go.importObject);
        go.run(result.instance);
    }

    // async process(byteArray: Uint8Array, size: number): Promise<Uint8Array> {
    //     return new Promise((resolve, reject) => {
    //         const threshold = 0.75;
    //         try {
    //             resolve((window as any).processImage(byteArray, threshold, size));
    //             return;
    //         } catch {
    //             reject();
    //         }
    //     });
    // }

    process(byteArray: Uint8Array, { threshold, size }: Settings): Uint8Array {
        return (window as any).processImage(byteArray, threshold, size);
    }
}

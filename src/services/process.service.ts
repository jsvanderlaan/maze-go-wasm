import { inject, Injectable } from '@angular/core';
import { BehaviorSubject, combineLatest, map, Observable, ReplaySubject, Subject } from 'rxjs';
import { Settings, SettingsComponent } from 'src/app/settings.component';
import { WasmService } from './wasm.service';

@Injectable({
    providedIn: 'root',
})
export class ProcessService {
    private wasmService = inject(WasmService);
    readonly original: Subject<Uint8Array> = new ReplaySubject(1);
    readonly settings: Subject<Settings> = new BehaviorSubject(SettingsComponent.defaultSettings);

    readonly output: Observable<Uint8Array> = combineLatest([this.original, this.settings]).pipe(
        map(([original, settings]) => this.wasmService.process(original, settings.size))
    );
}

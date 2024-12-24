import { inject, Injectable } from '@angular/core';
import { BehaviorSubject, combineLatest, from, merge, Observable, ReplaySubject, Subject, switchMap } from 'rxjs';
import { Settings, SettingsComponent } from 'src/app/settings.component';
import { inspectStatus } from 'src/helpers/inspectStatus';
import { Status } from 'src/types/status.type';
import { WorkerService } from './worker.service';

@Injectable({
    providedIn: 'root',
})
export class ProcessService {
    private workerService = inject(WorkerService);

    readonly sourceImage: Subject<Uint8Array> = new ReplaySubject(1);
    readonly sourceText: Subject<{ text: string; outline: boolean }> = new ReplaySubject(1);
    readonly settings: Subject<Settings> = new BehaviorSubject(SettingsComponent.defaultSettings);

    private readonly _processedImage = combineLatest([this.sourceImage, this.settings]).pipe(
        switchMap(([original, settings]) => from(this.workerService.processImage(original, settings)).pipe(inspectStatus))
    );

    private readonly _processedText = combineLatest([this.sourceText, this.settings]).pipe(
        switchMap(([{ text, outline }, settings]) =>
            from(this.workerService.processText(text, outline, settings)).pipe(inspectStatus)
        )
    );

    readonly output: Observable<Status<Uint8Array>> = merge(this._processedImage, this._processedText);
}

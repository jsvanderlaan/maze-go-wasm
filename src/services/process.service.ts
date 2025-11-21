import { inject, Injectable } from '@angular/core';
import { from, merge, Observable, ReplaySubject, Subject, switchMap } from 'rxjs';
import { inspectStatus } from 'src/helpers/inspectStatus';
import { ImageSourceInput, TextSourceInput } from 'src/types/setting.type';
import { Status } from 'src/types/status.type';
import { WorkerService } from './worker.service';

@Injectable({
    providedIn: 'root',
})
export class ProcessService {
    private workerService = inject(WorkerService);

    readonly sourceImage: Subject<ImageSourceInput> = new ReplaySubject(1);
    readonly sourceText: Subject<TextSourceInput> = new ReplaySubject(1);

    private readonly _processedImage = this.sourceImage.pipe(
        switchMap(input => from(this.workerService.processImage(input)).pipe(inspectStatus))
    );

    private readonly _processedText = this.sourceText.pipe(
        switchMap(input => from(this.workerService.processText(input)).pipe(inspectStatus))
    );

    readonly output: Observable<Status<Uint8Array>> = merge(this._processedImage, this._processedText);
}

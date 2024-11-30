import { CommonModule } from '@angular/common';
import { Component, computed, input } from '@angular/core';
import { ByteArrayHelper } from 'src/helpers/byte-array.helper';

@Component({
    imports: [CommonModule],
    selector: 'app-result',
    templateUrl: './result.component.html',
})
export class ResultComponent {
    static readonly mimeTypeResult = 'image/png';
    readonly result = input.required<Uint8Array>();
    readonly url = computed(() => ByteArrayHelper.toUrl(this.result()));

    readonly shareData = computed(() => ({
        files: [
            new File([ByteArrayHelper.toBlob(this.result(), ResultComponent.mimeTypeResult)], ResultComponent.fileName(), {
                type: ResultComponent.mimeTypeResult,
                lastModified: new Date().getTime(),
            }),
        ],
    }));

    readonly canShare = computed(() => navigator.canShare && navigator.canShare(this.shareData()));

    async share(): Promise<void> {
        if (this.canShare()) {
            await navigator.share(this.shareData());
        }
    }

    async copy(): Promise<void> {
        const blob = ByteArrayHelper.toBlob(this.result(), ResultComponent.mimeTypeResult);
        const data = [new ClipboardItem({ [ResultComponent.mimeTypeResult]: blob })];
        await navigator.clipboard.write(data);
    }

    async download(): Promise<void> {
        const blob = ByteArrayHelper.toBlob(this.result(), ResultComponent.mimeTypeResult);
        const img = URL.createObjectURL(blob);

        const link = document.createElement('a');
        link.href = img;
        link.download = ResultComponent.fileName();
        document.body.appendChild(link);
        link.click();
        document.body.removeChild(link);
    }

    static fileName(): string {
        return `a_maze_ing-${new Date().toISOString()}.png`;
    }
}

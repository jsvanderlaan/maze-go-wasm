import { Component, EventEmitter, Output } from '@angular/core';

@Component({
    imports: [],
    selector: 'app-example-upload',
    template: `
        <div class="w-full bg-gray-900 text-gray-200 rounded shadow-lg">
            <div class="relative">
                <div
                    class="flex overflow-x-auto scrollbar-hide snap-x snap-mandatory space-x-4 p-4 bg-gray-800 rounded"
                    id="carousel"
                >
                    @for (image of examples; track image) {
                        <img
                            [src]="image"
                            alt="Example Maze"
                            class="max-w-full h-32 object-cover rounded cursor-pointer hover:ring-2 hover:ring-green-500 transition-all snap-center"
                            (click)="selectImage(image)"
                        />
                    }
                </div>

                <button
                    class="absolute top-1/2 left-0 transform -translate-y-1/2 bg-gray-700/50 hover:bg-gray-600/50 rounded-full p-2"
                    (click)="scrollCarousel('left')"
                >
                    <svg xmlns="http://www.w3.org/2000/svg" class="w-6 h-6 text-gray-300" viewBox="0 0 20 20" fill="currentColor">
                        <path
                            fill-rule="evenodd"
                            d="M12.293 14.707a1 1 0 010-1.414L8.414 10l3.879-3.879a1 1 0 10-1.414-1.414l-5 5a1 1 0 000 1.414l5 5a1 1 0 001.414 0z"
                            clip-rule="evenodd"
                        />
                    </svg>
                </button>

                <button
                    class="absolute top-1/2 right-0 transform -translate-y-1/2 bg-gray-700/50 hover:bg-gray-600/50 rounded-full p-2"
                    (click)="scrollCarousel('right')"
                >
                    <svg xmlns="http://www.w3.org/2000/svg" class="w-6 h-6 text-gray-300" viewBox="0 0 20 20" fill="currentColor">
                        <path
                            fill-rule="evenodd"
                            d="M7.707 14.707a1 1 0 000-1.414L11.586 10 7.707 6.121a1 1 0 111.414-1.414l5 5a1 1 0 010 1.414l-5 5a1 1 0 01-1.414 0z"
                            clip-rule="evenodd"
                        />
                    </svg>
                </button>
            </div>
        </div>
    `,
})
export class ExampleUploadComponent {
    @Output() next = new EventEmitter<Uint8Array>();

    readonly examples = ['balloon.jpg', 'pikachu.jpg', 'heart.jpg', 'square.png', 'go.png', 'infinity.jpg', 'donut.png'].map(
        x => `assets/examples/${x}`
    );

    async selectImage(image: string): Promise<void> {
        const response = await fetch(image);
        const array = new Uint8Array(await response.arrayBuffer());
        this.next.emit(array);
    }

    scrollCarousel(direction: 'left' | 'right'): void {
        const carousel = document.getElementById('carousel');
        if (carousel) {
            const scrollAmount = direction === 'left' ? -32 : 32;
            carousel.scrollBy({ left: scrollAmount, behavior: 'smooth' });
        }
    }
}

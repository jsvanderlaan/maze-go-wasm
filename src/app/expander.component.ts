import { NgClass } from '@angular/common';
import { Component, Input, input, OnInit } from '@angular/core';
import { Observable } from 'rxjs';

@Component({
    imports: [NgClass],
    selector: 'app-expander',
    template: `
        <button
            class="w-full flex justify-between items-center p-2 px-4 bg-gray-800 rounded hover:bg-gray-700 transition-all"
            (click)="show = !show"
        >
            <span class="text-md font-small">{{ header() }}</span>
            <svg
                class="w-5 h-5 transition-transform"
                [ngClass]="{ 'rotate-180': show }"
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
            >
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
            </svg>
        </button>
        @if (show) {
            <div class="mt-2 p-4 bg-gray-800 rounded transition-all ease-in-out duration-300">
                <ng-content />
            </div>
        }
    `,
})
export class ExpanderComponent implements OnInit {
    @Input({ required: true }) toggle!: Observable<any>;
    readonly header = input.required<string>();
    show = false;

    ngOnInit(): void {
        this.toggle.subscribe(() => (this.show = !this.show));
        this.show = false;
    }
}

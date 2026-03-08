import type { CollectionQueryBuilder, PageCollectionItemBase } from '~/src/types';
type Section = {
    id: string;
    title: string;
    titles: string[];
    level: number;
    content: string;
};
export type GenerateSearchSectionsOptions = {
    ignoredTags?: string[];
    extraFields?: (string | symbol | number)[];
    minHeading?: `h${1 | 2 | 3 | 4 | 5 | 6}`;
    maxHeading?: `h${1 | 2 | 3 | 4 | 5 | 6}`;
};
export declare function generateSearchSections<T extends PageCollectionItemBase>(queryBuilder: CollectionQueryBuilder<T>, opts?: GenerateSearchSectionsOptions): Promise<Section[]>;
export {};

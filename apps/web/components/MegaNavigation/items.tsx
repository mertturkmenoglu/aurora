export interface L2Item {
    id: string;
    name: string;
}

export interface L1Item {
    id: string;
    name: string;
    items: L2Item[];
}

export interface L0Item {
    id: string;
    name: string;
    items: L1Item[];
}

export interface MegaNavigationItems {
    items: L0Item[];
}

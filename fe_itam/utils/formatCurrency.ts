export const formatCurrency = (value: number | undefined | null): string => {
    if (value === undefined || value === null) return '';
    return new Intl.NumberFormat('id-ID', {
        style: 'currency',
        currency: 'IDR',
        minimumFractionDigits: 0,
        maximumFractionDigits: 0
    }).format(value);
};

export const parseCurrency = (value: string): number => {
    if (!value) return 0;
    // Remove currency symbol, thousand separators, and other non-numeric characters
    return Number(value.replace(/[^\d,-]/g, ''));
};
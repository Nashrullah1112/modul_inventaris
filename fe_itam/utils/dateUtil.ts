// src/utils/dateUtils.ts

/**
 * Parses a date string into a formatted date
 * @param dateString - The input date string
 * @param format - Optional format (defaults to 'YYYY-MM-DD')
 * @returns Formatted date string
 */
export function parseDate(
    dateString: string | Date,
    format: "YYYY-MM-DD" | "DD/MM/YYYY" | "MM/DD/YYYY" | "FULL-DD/MM/YYYY" = "YYYY-MM-DD"
): string {
    // If already a Date object, convert to string
    const inputDate = dateString instanceof Date ? dateString : new Date(dateString);

    // Handle invalid dates
    if (isNaN(inputDate.getTime())) {
        console.warn(`Invalid date: ${dateString}`);
        return "Invalid Date";
    }

    // Day names in Indonesian
    const dayNames = [
        "Minggu", "Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu"
    ];

    // Month names in Indonesian
    const monthNames = [
        "Januari", "Februari", "Maret", "April", "Mei", "Juni",
        "Juli", "Agustus", "September", "Oktober", "November", "Desember"
    ];

    // Extract date components
    const year = inputDate.getFullYear();
    const monthIndex = inputDate.getMonth(); // 0-based index
    const month = String(monthIndex + 1).padStart(2, '0');  // Numeric month (01, 02, etc.)
    const monthName = monthNames[monthIndex]; // Get the month name in Indonesian
    const day = String(inputDate.getDate()).padStart(2, '0');
    const dayName = dayNames[inputDate.getDay()];

    // Return formatted date based on specified format
    switch (format) {
        case "YYYY-MM-DD":
            return `${year}-${month}-${day}`;
        case "DD/MM/YYYY":
            return `${day}/${month}/${year}`;
        case "MM/DD/YYYY":
            return `${month}/${day}/${year}`;
        case "FULL-DD/MM/YYYY":
            return `${dayName}, ${day} ${monthName} ${year}`;
        default:
            return `${year}-${month}-${day}`;  // Default format if not recognized
    }
}


/**
 * Formats a date with time in a localized, readable format
 * @param dateString - The input date string
 * @returns Formatted date and time string
 */
export function formatDateTime(dateString: string | Date): string {
    const inputDate = dateString instanceof Date
        ? dateString
        : new Date(dateString);

    // Handle invalid dates
    if (isNaN(inputDate.getTime())) {
        console.warn(`Invalid date: ${dateString}`);
        return 'Invalid Date';
    }

    // Use toLocaleString for a human-readable format
    return inputDate.toLocaleString('id-ID', {
        year: 'numeric',
        month: 'long',
        day: 'numeric',
        hour: '2-digit',
        minute: '2-digit',
        second: '2-digit',
        hour12: false
    });
}

/**
 * Calculate time difference between two dates
 * @param startDate - The start date
 * @param endDate - The end date (defaults to current date)
 * @returns Object with time differences
 */
export function getDateDifference(startDate: string | Date, endDate?: string | Date) {
    const start = startDate instanceof Date
        ? startDate
        : new Date(startDate);

    const end = endDate
        ? (endDate instanceof Date ? endDate : new Date(endDate))
        : new Date();

    // Handle invalid dates
    if (isNaN(start.getTime()) || isNaN(end.getTime())) {
        console.warn(`Invalid date for comparison: ${startDate}, ${endDate}`);
        return null;
    }

    const diffMs = Math.abs(end.getTime() - start.getTime());
    const diffDays = Math.floor(diffMs / (1000 * 60 * 60 * 24));
    const diffYears = end.getFullYear() - start.getFullYear();
    const diffMonths = end.getMonth() - start.getMonth();

    return {
        totalDays: diffDays,
        years: diffYears,
        months: Math.abs(diffMonths),
        fullYears: diffYears + (diffMonths >= 0 ? 0 : -1)
    };
}

/**
 * Check if a date is today
 * @param dateString - The input date string
 * @returns Boolean indicating if the date is today
 */
export function isToday(dateString: string | Date): boolean {
    const inputDate = dateString instanceof Date
        ? dateString
        : new Date(dateString);

    const today = new Date();

    return (
        inputDate.getDate() === today.getDate() &&
        inputDate.getMonth() === today.getMonth() &&
        inputDate.getFullYear() === today.getFullYear()
    );
}
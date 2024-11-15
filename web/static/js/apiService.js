export async function scrapeUrl(url, prompt, responseStructure) {
    const response = await fetch('/api/v1/scrape', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ url, prompt, responseStructure })
    });

    if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
    }

    const data = await response.json();
    return data.results;
}

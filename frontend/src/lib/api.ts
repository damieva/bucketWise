const BASE_URL = process.env.NEXT_PUBLIC_API_URL;

export async function api<T>(path: string, options?: RequestInit): Promise<T> {
    const res = await fetch(`${BASE_URL}${path}`, {
        ...options,
        headers: {
            "Content-Type": "application/json",
            ...(options?.headers ?? {})
        }
    });

    if (!res.ok) {
        throw new Error(await res.text());
    }

    return res.json();
}

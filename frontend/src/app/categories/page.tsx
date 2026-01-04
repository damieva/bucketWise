import { api } from "@/lib/api";
import { Card, CardHeader, CardTitle, CardContent } from "@/components/ui/card";
import { Button } from "@/components/ui/button";

type Category = {
    id: string;
    name: string;
    type: string;
};

type CategoryResponse = {
    categories: Category[];
};

export default async function CategoriesPage() {
    const data = await api<CategoryResponse>("/categories");
    const categories = data.categories;

    return (
        <div className="w-full max-w-4xl mx-auto space-y-6">
            {/* Top bar */}
            <div className="flex items-center justify-between">
                <h1 className="text-3xl font-bold tracking-tight">Categories</h1>
                <Button>Add Category</Button>
            </div>

            {/* Categories card */}
            <Card className="bg-card text-card-foreground shadow-sm border border-border">
                <CardHeader>
                    <CardTitle className="text-xl">Existing Categories</CardTitle>
                </CardHeader>

                <CardContent>
                    {categories.length === 0 ? (
                        <p className="text-muted-foreground">No categories available.</p>
                    ) : (
                        <div className="divide-y border border-border rounded-md bg-card">
                            {categories.map((cat) => (
                                <div
                                    key={cat.id}
                                    className="grid grid-cols-2 p-4 hover:bg-muted transition"
                                >
                                    <span className="font-medium">{cat.name}</span>
                                    <span className="text-muted-foreground">{cat.type}</span>
                                </div>
                            ))}
                        </div>
                    )}
                </CardContent>
            </Card>
        </div>
    );
}

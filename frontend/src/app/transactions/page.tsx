import { api } from "@/lib/api";
import { Card, CardHeader, CardTitle, CardContent } from "@/components/ui/card";

type Transaction = {
    id: string;
    amount: number;
    date: string;
    description: string;
    category_id: string;
    category_name: string;
    type: string;
};

type TransactionResponse = {
    transactions: Transaction[];
};

export default async function TransactionsPage() {
    const data = await api<TransactionResponse>("/transactions");
    const transactions = data.transactions;

    return (
        <div className="w-full max-w-4xl mx-auto space-y-6">
            {/* Top bar */}
            <div className="flex items-center justify-between">
                <h1 className="text-3xl font-bold tracking-tight">Transactions</h1>
            </div>

            {/* Transactions card */}
            <Card className="bg-card text-card-foreground shadow-sm border border-border">
                <CardHeader>
                    <CardTitle className="text-xl">Transactions List</CardTitle>
                </CardHeader>

                <CardContent>
                    {transactions.length === 0 ? (
                        <p className="text-muted-foreground">No transactions available.</p>
                    ) : (
                        <div className="divide-y border border-border rounded-md bg-card">
                            {transactions.map((tx) => (
                                <div
                                    key={tx.id}
                                    className="grid grid-cols-4 gap-4 p-4 hover:bg-muted transition"
                                >
                                    <span className="font-medium">{tx.date}</span>
                                    <span className="text-muted-foreground">{tx.amount}</span>
                                    <span className="text-muted-foreground">{tx.type}</span>
                                    <span className="text-muted-foreground">{tx.category_name}</span>
                                </div>
                            ))}
                        </div>
                    )}
                </CardContent>
            </Card>
        </div>
    );
}

"use client";

import React, { use } from "react";
import { GoalItem, MobileView } from "@/components/GoalTracker";

interface PageProps {
  params: Promise<{ userId: string }>;
}

export default function TodayPage({ params }: PageProps) {
  const { userId } = use(params);

  // Generate 24 static goal slots
  const goals = Array.from({ length: 24 }, (_, i) => ({
    id: i,
    time: i < 10 ? `0${i}:00` : `${i}:00`,
  }));

  return (
    <MobileView>
      {/* Header: Indigo Background, White Text */}
      <header className="p-6 bg-indigo-600 text-white">
        <h1 className="text-lg font-bold tracking-tight">Goals this month:</h1>
        <p className="text-xs opacity-80 font-mono mt-1">UID: {userId}</p>
      </header>

      {/* Main List: White/Gray context */}
      <main className="p-4 bg-white">
        {goals.map((goal) => (
          <GoalItem key={goal.id} text={`${goal.time} - Planned Task`} />
        ))}
      </main>

      {/* Minimal Footer: Gray border, White background */}
      <footer className="p-6 border-t border-gray-100 text-center">
        <p className="text-[10px] text-gray-400 uppercase tracking-widest">
          End of List
        </p>
      </footer>
    </MobileView>
  );
}

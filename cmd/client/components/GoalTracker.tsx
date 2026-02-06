"use client";

import React, { useState } from "react";

/**
 * GOAL ITEM
 * Strict Palette: White (BG), Gray (Border/Text), Indigo (Action)
 */
export const GoalItem = ({ text }: { text: string }) => {
  const [isDone, setIsDone] = useState(false);

  return (
    <div
      onClick={() => setIsDone(!isDone)}
      className={`flex items-center p-4 mb-3 border rounded-lg cursor-pointer transition-all ${
        isDone ? "border-indigo-600 bg-gray-50" : "border-gray-200 bg-white"
      }`}
    >
      {/* Checkbox using Indigo for action, Gray for inactive */}
      <div
        className={`w-5 h-5 rounded border mr-4 flex items-center justify-center ${
          isDone ? "bg-indigo-600 border-indigo-600" : "border-gray-400"
        }`}
      >
        {isDone && <span className="text-white text-[10px]">✓</span>}
      </div>

      {/* Text using Gray tones */}
      <span
        className={`text-sm font-medium ${isDone ? "text-gray-400 line-through" : "text-gray-900"}`}
      >
        {text}
      </span>
    </div>
  );
};

/**
 * MOBILE VIEW
 * Basic Gray background with a White center column
 */
export const MobileView = ({ children }: { children: React.ReactNode }) => (
  <div className="min-h-screen bg-gray-100 flex justify-center">
    <div className="w-full max-w-md bg-white min-h-screen shadow-sm">
      {children}
    </div>
  </div>
);

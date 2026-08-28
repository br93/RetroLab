import React, { useState, useEffect } from 'react';

export function CompletedGamesHistory({ activeGameId, onSelectGame }) {
  const [history, setHistory] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetch('/api/retro/completed')
      .then((res) => res.json())
      .then((data) => {
        setHistory(Array.isArray(data) ? data : []);
        setLoading(false);
      })
      .catch((err) => {
        console.error("History data loading ledger problem:", err);
        setLoading(false);
      });
  }, []);

  if (loading) return <div className="text-slate-500 text-xs font-mono uppercase">Syncing mastery ledgers...</div>;

  const softcoreGames = history.filter((game) => game.HardcoreMode === "0");

  return (
    <div className="grid grid-cols-2 gap-3 max-h-[480px] overflow-y-auto pr-1">
      {softcoreGames.map((item) => {
        const isSelected = activeGameId === item.GameID;

        return (
          <div 
            key={`${item.GameID}-${item.HardcoreMode}`}
            onClick={() => onSelectGame(item.GameID, item.Title)}
            className={`p-3 rounded-xl flex items-center space-x-3 cursor-pointer transition-all border ${
              isSelected 
                ? 'bg-indigo-500/10 border-indigo-500 shadow-md ring-1 ring-indigo-500/20' 
                : 'bg-slate-950/40 border-slate-900 hover:border-slate-800'
            }`}
          >
            {/* Fixed template image formatting parameters mapping string arrays */}
            <img 
              src={`https://retroachievements.org${item.ImageIcon}`} 
              alt={item.Title} 
              className="w-10 h-10 object-contain bg-slate-900 rounded border border-slate-800/40 p-0.5" 
            />
            <div className="min-w-0 flex-1">
              <h4 className={`font-bold text-xs truncate ${isSelected ? 'text-indigo-400' : 'text-slate-200'}`}>
                {item.Title}
              </h4>
              <div className="flex justify-between items-center mt-1">
                <span className="text-[9px] text-slate-500 font-mono">{item.NumAwarded}/{item.MaxPossible}</span>
                <span className="inline-block text-[8px] font-black uppercase tracking-wider bg-emerald-500/10 text-emerald-400 px-1.5 py-0.5 rounded border border-emerald-500/20">
                  {item.PctWon === "1.0000" ? "Mastered" : "Active"}
                </span>
              </div>
            </div>
          </div>
        );
      })}
    </div>
  );
}
import React from 'react';

export function GameProgressionList({ recentUnlocks, activeGameId, onSelectGame }) {
  return (
    <div className="grid grid-cols-2 gap-3 max-h-[480px] overflow-y-auto pr-1">
      {recentUnlocks.map((unlock) => {
        const isSelected = activeGameId === unlock.GameID;

        return (
          <div
            key={unlock.AchievementID}
            onClick={() => onSelectGame(unlock.GameID, unlock.GameTitle)}
            className={`p-3 rounded-xl flex items-center space-x-3 cursor-pointer transition-all duration-200 border ${
              isSelected
                ? 'border-indigo-500 bg-indigo-500/[0.03] shadow-md ring-1 ring-indigo-500/20'
                : 'bg-slate-950/60 border-slate-900 hover:border-slate-800'
            }`}
          >
            {/* Fixed Image component template interpolation tag link string mapping layout */}
            <img 
              src={`https://retroachievements.org${unlock.BadgeURL}`} 
              alt={unlock.Title} 
              className="w-10 h-10 object-contain rounded bg-slate-900 p-0.5 shadow border border-slate-800/40" 
            />
            <div className="min-w-0 flex-1">
              <h4 className={`font-bold text-xs truncate ${isSelected ? 'text-indigo-400' : 'text-amber-400'}`}>
                {unlock.Title}
              </h4>
              <p className="text-[10px] text-slate-400 truncate mt-0.5">{unlock.GameTitle}</p>
              <div className="text-[8px] font-mono text-slate-600 mt-1 uppercase tracking-wider">
                {unlock.Date ? new Date(unlock.Date).toLocaleDateString() : 'Just Now'}
              </div>
            </div>
          </div>
        );
      })}
    </div>
  );
}
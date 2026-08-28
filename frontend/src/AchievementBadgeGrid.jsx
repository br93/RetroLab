import React from 'react';

export function AchievementBadgeGrid({ gameId, gameTitle, recentUnlocks }) {
  if (!gameId) {
    return (
      <div className="h-full min-h-[500px] flex flex-col items-center justify-center text-center p-6 border-2 border-dashed border-slate-900 rounded-xl">
        <span className="text-xl mb-2">🎮</span>
        <p className="text-xs text-slate-500 max-w-xs font-mono uppercase tracking-wider">
          Select an entry item from the panels tab window to inspect achievements dashboard matrices.
        </p>
      </div>
    );
  }

  const filteredBadges = recentUnlocks.filter((unlock) => unlock.GameID === gameId);

  return (
    <div className="space-y-4 h-full flex flex-col">
      <div className="border-b border-slate-900 pb-4">
        <span className="text-[9px] font-black text-indigo-400 uppercase tracking-widest font-mono">INSIDERS PROFILE VAULT</span>
        <h3 className="text-sm font-black text-white uppercase tracking-wider mt-1 truncate max-w-full">{gameTitle}</h3>
      </div>

      <div className="flex-1 overflow-y-auto max-h-[460px] pr-1">
        {!filteredBadges.length ? (
          <div className="text-slate-600 font-mono text-xs uppercase p-4">No milestone records inside lookback logs filter.</div>
        ) : (
          <div className="grid grid-cols-3 sm:grid-cols-4 gap-2.5">
            {filteredBadges.map((ach) => (
              <div 
                key={ach.AchievementID} 
                className="group relative flex flex-col items-center p-2 rounded-xl bg-slate-950/80 border border-slate-900 transition-all duration-300 hover:border-amber-500/40"
              >
                {/* Fixed Template Interp string concatenation parsing tag */}
                <img 
                  src={`https://retroachievements.org${ach.BadgeURL}`} 
                  alt={ach.Title} 
                  className="w-11 h-11 object-contain rounded ring-1 ring-slate-800/60 shadow-md group-hover:scale-105 transition-transform" 
                />
                
                <span className="text-[9px] font-mono text-amber-500/80 font-bold mt-1.5 bg-amber-500/5 border border-amber-500/10 px-1.5 py-0.5 rounded truncate max-w-full">
                  {ach.Points} PTS
                </span>

                {/* OVERLAY POPUP FLOATER PANEL */}
                <div className="absolute bottom-full mb-2 hidden group-hover:flex flex-col w-48 p-2.5 bg-[#0a0d16] text-[11px] rounded-lg border border-slate-900 shadow-2xl z-50 pointer-events-none left-1/2 -translate-x-1/2">
                  <span className="font-bold text-amber-400 truncate">{ach.Title}</span>
                  <span className="text-slate-400 text-[10px] mt-1 leading-normal">{ach.Description}</span>
                  <span className="text-[9px] text-slate-500 font-mono mt-1.5 pt-1 border-t border-slate-900">
                    🏆 {ach.Date ? new Date(ach.Date).toLocaleDateString() : 'Earned'}
                  </span>
                </div>
              </div>
            ))}
          </div>
        )}
      </div>
    </div>
  );
}
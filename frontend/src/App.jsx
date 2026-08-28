import React, { useState, useEffect } from 'react';
import { GameProgressionList } from './GameProgressionList';
import { CompletedGamesHistory } from './CompletedGamesHistory';
import { AchievementBadgeGrid } from './AchievementBadgeGrid';
import { LeftGamePanel } from './LeftGamePanel';

export default function App() {
  const [playerData, setPlayerData] = useState(null);
  const [recentUnlocks, setRecentUnlocks] = useState([]);
  const [minutesLookBack, setMinutesLookBack] = useState("43200"); 
  const [selectedGame, setSelectedGame] = useState({ id: null, title: "" });
  const [activeTab, setActiveTab] = useState('recent'); // Toggles right panel views
  const [loading, setLoading] = useState(true);
  const [panelOpen, setPanelOpen] = useState(false);
  const [panelDetails, setPanelDetails] = useState(null);
  const [panelLoading, setPanelLoading] = useState(false);

  useEffect(() => {
    fetch('/api/retro/player?g=3&a=5')
      .then((res) => res.json())
      .then((data) => setPlayerData(data))
      .catch((err) => console.error("Summary payload loading issue:", err));
  }, []);

  useEffect(() => {
    fetch(`/api/retro/recent?m=${minutesLookBack}`)
      .then((res) => res.json())
      .then((data) => {
        setRecentUnlocks(Array.isArray(data) ? data : []);
        setLoading(false);
      })
      .catch((err) => {
        console.error("Recent log ticker collection tracking issue:", err);
        setLoading(false);
      });
  }, [minutesLookBack]);

  const handleGameSelectFromArchive = (id, title) => {
    setSelectedGame({ id, title });
    setPanelOpen(true);
    setPanelLoading(true);

    fetch(`/api/retro/details?id=${id}`)
      .then((res) => res.json())
      .then((data) => {
        setPanelDetails(data);
        setPanelLoading(false);
      })
      .catch((err) => {
        console.error("Failed gathering progression matrix paths:", err);
        setPanelLoading(false);
      });
  };

  if (loading) {
    return <div className="min-h-screen bg-[#070a13] flex items-center justify-center text-indigo-400 font-mono text-xs tracking-widest">LOADING RETRO CORRIDORS //</div>;
  }

  return (
    <div className="min-h-screen bg-[#070a13] text-slate-100 font-sans antialiased bg-[radial-gradient(ellipse_at_top,_var(--tw-gradient-stops))] from-slate-900/40 via-[#070a13] to-[#03050a]">
      
      {/* 👇 MOUNT MOUNTED OVERLAY SEPARATELY AT THE CANVAS TOP-LEVEL STACK */}
      <LeftGamePanel isOpen={panelOpen} onClose={() => setPanelOpen(false)} details={panelDetails} loading={panelLoading} />

      <header className="border-b border-slate-900 bg-slate-950/40 backdrop-blur-md sticky top-0 z-40">
        <div className="max-w-7xl mx-auto px-4 h-16 flex items-center justify-between">
          <div className="flex items-center space-x-3">
            <div className="w-2.5 h-2.5 bg-indigo-500 rounded-sm animate-pulse shadow-[0_0_10px_#6366f1]"></div>
            <h1 className="text-xs font-black tracking-widest text-slate-300 uppercase">RETRO<span className="text-indigo-400">LAB</span> // HUB</h1>
          </div>
          {playerData && (
            <div className="flex items-center space-x-4">
              <span className="hidden md:inline text-[11px] text-slate-400 font-mono bg-slate-950/60 border border-slate-900 px-3 py-1 rounded">🎮 {playerData.RichPresenceMsg}</span>
              <div className="flex items-center space-x-3 bg-slate-950/60 border border-slate-900 px-3 py-1.5 rounded-full text-xs">
                <img src={`https://retroachievements.org${playerData.UserPic}`} alt={playerData.User} className="w-5 h-5 rounded-full ring-1 ring-indigo-500/30" />
                <span className="font-semibold text-slate-300">{playerData.User}</span>
                <span className="bg-indigo-500/10 text-indigo-400 border border-indigo-500/20 px-2 py-0.5 rounded-full font-bold text-[11px]">🏆 {playerData.TotalPoints}</span>
              </div>
            </div>
          )}
        </div>
      </header>

      <main className="max-w-7xl mx-auto px-4 py-8">
        <div className="grid lg:grid-cols-12 gap-8 items-start">
          
          <div className="lg:col-span-5 bg-slate-950/40 border border-slate-900 rounded-2xl p-6 shadow-2xl backdrop-blur-md min-h-[580px]">
            <AchievementBadgeGrid gameId={selectedGame.id} gameTitle={selectedGame.title} recentUnlocks={recentUnlocks} />
          </div>

          <div className="lg:col-span-7 bg-slate-950/40 border border-slate-900 rounded-2xl p-6 shadow-2xl backdrop-blur-md min-h-[580px] flex flex-col">
            <div className="flex items-center justify-between pb-4 border-b border-slate-900 mb-6">
              <div className="flex space-x-2 bg-slate-950/80 p-1 border border-slate-900 rounded-xl">
                <button onClick={() => setActiveTab('recent')} className={`px-4 py-2 text-xs font-bold uppercase tracking-wider rounded-lg transition-all ${activeTab === 'recent' ? 'bg-indigo-500/10 text-indigo-400 border border-indigo-500/20' : 'text-slate-400 hover:text-slate-200'}`}>Unlocks Log</button>
                <button onClick={() => setActiveTab('completed')} className={`px-4 py-2 text-xs font-bold uppercase tracking-wider rounded-lg transition-all ${activeTab === 'completed' ? 'bg-indigo-500/10 text-indigo-400 border border-indigo-500/20' : 'text-slate-400 hover:text-slate-200'}`}>Mastery Archive</button>
              </div>

              {activeTab === 'recent' && (
                <div className="flex items-center space-x-2 text-[10px] text-slate-500">
                  <span className="font-mono">WINDOW:</span>
                  <select value={minutesLookBack} onChange={(e) => setMinutesLookBack(e.target.value)} className="bg-slate-950 border border-slate-900 text-slate-400 rounded px-2 py-0.5 outline-none cursor-pointer font-mono">
                    <option value="1440">24H</option>
                    <option value="10080">7D</option>
                    <option value="43200">30D</option>
                  </select>
                </div>
              )}
            </div>

            <div className="flex-1">
              {activeTab === 'recent' ? (
                <GameProgressionList recentUnlocks={recentUnlocks} activeGameId={selectedGame.id} onSelectGame={(id, title) => setSelectedGame({ id, title })} />
              ) : (
                <CompletedGamesHistory activeGameId={selectedGame.id} onSelectGame={handleGameSelectFromArchive} /> // 👇 CONNECT PASSTHROUGH TO MODAL HOOK OVERLAY
              )}
            </div>

          </div>
        </div>
      </main>
    </div>
  );
}

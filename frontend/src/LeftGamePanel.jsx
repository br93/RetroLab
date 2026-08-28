import React, { useState, useRef } from 'react';

export function LeftGamePanel({ isOpen, onClose, details, loading }) {
  if (!isOpen) return null;

  const fileInputRef = useRef(null);
  const [uploading, setUploading] = useState(false);
  const [uploadStatus, setUploadStatus] = useState('');

  // Render a loading skeleton while fetching deep data from the backend proxy
  if (loading || !details) {
    return (
      <div className="fixed top-0 left-0 h-full w-80 sm:w-96 bg-[#090d16] border-r border-slate-900 shadow-2xl z-50 p-6 font-mono text-[11px] text-indigo-400 flex flex-col justify-center items-center">
        <div className="animate-spin rounded-full h-4 w-4 border-2 border-indigo-500 border-t-transparent mb-2"></div>
        DECRYPTING GAME PROFILES...
      </div>
    );
  }

  // Turn achievements dictionary object collection into flat array entries
  const achList = details.Achievements ? Object.values(details.Achievements) : [];

  // Total count of all achievements configured for this retro game title
  const totalAchievementsCount = achList.length;

  const gameBannerUrl = details.ImageTitle?.includes('000002.png')
    ? details.ImageIcon
    : details.ImageTitle;

  // Trigger the hidden native file input browser selection dialog
  const handleUploadButtonClick = () => {
    fileInputRef.current?.click();
  };

  // POST Multi-part streaming pipeline to submit save file binaries to Go container storage volumes
  const handleFileChange = async (e) => {
    const selectedFile = e.target.files?.[0];
    if (!selectedFile) return;

    setUploading(true);
    setUploadStatus('STREAMING...');

    const formData = new FormData();
    formData.append('saveState', selectedFile);

    try {
      const response = await fetch(`/api/v1/retro/upload/savestate?id=${details.ID}`, {
        method: 'POST',
        body: formData,
      });

      if (!response.ok) {
        throw new Error('Server storage target registration issue.');
      }

      setUploadStatus('✅ SUCCESS!');
      setTimeout(() => setUploadStatus(''), 3000);
    } catch (err) {
      console.error(err);
      setUploadStatus('❌ FAILED');
      setTimeout(() => setUploadStatus(''), 4000);
    } finally {
      setUploading(false);
      if (fileInputRef.current) fileInputRef.current.value = ''; // Reset file input buffer
    }
  };

  return (
    <>
      {/* BACKGROUND CLICK HANDLER OVERLAY BLOCK */}
      <div className="fixed inset-0 bg-black/70 backdrop-blur-sm z-40 transition-all" onClick={onClose} />

      {/* LEFT ACCENT DRAWER FRAME CONTAINER */}
      <div className="fixed top-0 left-0 h-full w-80 sm:w-96 bg-[#080b14] border-r border-slate-900 shadow-2xl z-50 p-6 flex flex-col justify-between font-sans animate-[slideIn_0.25s_ease-out]">

        {/* UPPER PANEL CONTENT BLOCK */}
        <div className="flex flex-col min-h-0 flex-grow">
          {/* HEADER METRICS DESCRIPTORS */}
          <div className="flex items-center justify-between border-b border-slate-900 pb-4 mb-5 flex-shrink-0">
            <div className="min-w-0 flex-1">
              <span className="text-[9px] font-black text-indigo-400 uppercase tracking-widest font-mono">ARCHIVE SYSTEM PANEL</span>
              <h2 className="text-sm font-black text-white uppercase tracking-wider mt-0.5 truncate pr-2">
                {details.Title}
              </h2>
            </div>
            <button
              onClick={onClose}
              className="text-slate-500 hover:text-slate-300 font-mono text-xs bg-slate-950 border border-slate-900 px-2 py-1 rounded-lg transition"
            >
              ✕
            </button>
          </div>

          {/* MAIN GRAPHIC CONSOLE FRAME HERO IMAGE */}
          <div className="flex flex-col space-y-2 mb-5 flex-shrink-0">
            <div className="rounded-xl overflow-hidden border border-slate-900 bg-slate-950/80 p-3 flex items-center justify-center w-full aspect-square max-h-[280px]">
              <img
                src={`https://retroachievements.org${gameBannerUrl}`}
                alt={details.Title}
                className="w-full h-full object-contain transition-transform duration-300 hover:scale-102"
              />
            </div>
            <div className="self-start bg-slate-950 border border-slate-900 text-[9px] font-mono text-slate-400 px-2.5 py-1 rounded font-black tracking-wider uppercase">
              {details.ConsoleName}
            </div>
          </div>

          {/* TOTAL ACHIEVEMENTS DISPLAY SECTION */}
          <div className="mb-5 font-mono flex-shrink-0">
            <div className="flex items-center justify-between p-3 rounded-xl bg-indigo-500/5 border border-indigo-500/10">
              <span className="text-[10px] font-black text-slate-400 uppercase tracking-wider">ACHIEVEMENTS</span>
              <span className="text-xs font-black text-indigo-400 bg-indigo-500/10 border border-indigo-500/20 px-2 py-0.5 rounded">
                {details.UserCompletion}
              </span>
            </div>
          </div>

          {/* DYNAMIC SCROLL CONTAINER - EXPANDS TO ABSORB ALL EMPTY SPACE */}
          <div className="flex-1 min-h-0 flex flex-col mb-4">
            <h3 className="text-[10px] font-black font-mono text-slate-500 uppercase tracking-wider mb-2.5 flex-shrink-0">
              TOTAL ACHIEVEMENTS - {totalAchievementsCount}
            </h3>
            {achList.length > 0 ? (
              <div className="overflow-y-auto pr-1 flex-grow space-y-1.5 custom-scrollbar">
                {achList.map((ach) => {
                  // 👇 DETERMINE IF THE INDIVIDUAL MILESTONE HAS BEEN UNLOCKED
                  const isUnlocked = ach.DateEarned != "";

                  // 👇 IF LOCKED, APPEND THE OFFICIAL '_lock.png' SUFFIX
                  const badgeFilename = isUnlocked
                    ? `${ach.BadgeName}.png`
                    : `${ach.BadgeName}_lock.png`;

                  return (
                    <a
                      key={ach.ID}
                      href={`https://retroachievements.org/achievement/${ach.ID}`}
                      target="_blank"
                      rel="noopener noreferrer"
                      className="flex items-center space-x-2.5 p-1.5 rounded-lg bg-slate-950/30 border border-slate-900/60 transition cursor-pointer hover:border-indigo-500/50 hover:bg-slate-950/60 group"
                    >
                      <img
                        src={`https://retroachievements.org/Badge/${badgeFilename}`}
                        alt=""
                        className={`w-6 h-6 object-contain rounded bg-slate-950 border border-slate-800/40 transition group-hover:scale-105 ${isUnlocked ? 'opacity-90 group-hover:opacity-100' : 'opacity-40'
                          }`}
                      />
                      <span className={`text-xs font-sans truncate transition ${isUnlocked
                        ? 'text-slate-200 group-hover:text-indigo-400'
                        : 'text-slate-500 group-hover:text-slate-400 font-medium'
                        }`}>
                        {ach.Title}
                      </span>
                    </a>
                  );
                })}
              </div>
            ) : (
              <div className="text-slate-600 font-mono text-[10px] uppercase pl-1 flex-shrink-0">
                No active milestone registers logged.
              </div>
            )}
          </div>
        </div>

        {/* LOWER CONTROLS BLOCK - ANCHORED PERMANENTLY TO THE BOTTOM */}
        <div className="space-y-2.5 pt-4 border-t border-slate-900 flex-shrink-0 bg-[#080b14]">

          {/* Row 1: Full-Width ROM Downloader */}
          <a
            href={`/api/v1/retro/download/rom?id=${details.ID}`}
            download
            className="block text-center font-mono text-[9px] font-black tracking-widest bg-indigo-600/10 hover:bg-indigo-600/20 text-indigo-400 border border-indigo-500/20 py-3 rounded-xl transition w-full"
          >
            📦 DOWNLOAD ROM FILE
          </a>

          {/* Row 2: Split Action Columns */}
          <div className="grid grid-cols-2 gap-2.5">
            {/* Left Box: Save State Upload Trigger */}
            <div>
              <input
                type="file"
                ref={fileInputRef}
                onChange={handleFileChange}
                className="hidden"
                accept=".state,.srm,.sav,.dat"
              />
              <button
                onClick={handleUploadButtonClick}
                disabled={uploading}
                className={`w-full text-center font-mono text-[9px] font-black tracking-widest py-3 rounded-xl border transition duration-150 ${uploading
                  ? 'bg-amber-600/20 border-amber-500/30 text-amber-400 cursor-not-allowed'
                  : 'bg-slate-950 hover:bg-slate-900 border-slate-800 text-slate-400 hover:text-slate-200'
                  }`}
              >
                {uploadStatus || '📤 SAVE STATE'}
              </button>
            </div>

            {/* Right Box: Load State Stream Action Link */}
            <a
              href={`/api/v1/retro/download/savestate?id=${details.ID}`}
              download
              className="text-center flex items-center justify-center font-mono text-[9px] font-black tracking-widest bg-emerald-600/10 hover:bg-emerald-600/20 text-emerald-400 border border-emerald-500/20 py-3 rounded-xl transition"
            >
              📥 LOAD STATE
            </a>
          </div>

        </div>

      </div>
    </>
  );
}
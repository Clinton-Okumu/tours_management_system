import React from "react";

type ModalProps = {
  title: string;
  isOpen: boolean;
  onClose: () => void;
  children: React.ReactNode;
};

const Modal = ({ title, isOpen, onClose, children }: ModalProps) => {
  if (!isOpen) return null;

  return (
    <div className="fixed inset-0 z-50 bg-white/20 backdrop-blur-sm flex items-center justify-center px-4">
      <div className="relative bg-white w-full max-w-md rounded-xl shadow-lg p-6 animate-fade-in-up">
        {/* Close button */}
        <button
          onClick={onClose}
          className="absolute top-4 right-4 text-gray-500 hover:text-red-500 text-xl font-bold"
          aria-label="Close"
        >
          ×
        </button>

        {/* Title */}
        <h2 className="text-2xl font-bold text-gray-800 mb-4 text-center">
          {title}
        </h2>

        {/* Content (form goes here) */}
        <div>{children}</div>
      </div>
    </div>
  );
};

export default Modal;

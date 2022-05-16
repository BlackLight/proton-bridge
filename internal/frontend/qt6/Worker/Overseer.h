// Copyright (c) 2022 Proton AG
//
// This file is part of Proton Mail Bridge.
//
// Proton Mail Bridge is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Proton Mail Bridge is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Proton Mail Bridge. If not, see <https://www.gnu.org/licenses/>.


#ifndef BRIDGE_QT6_OVERSEER_H
#define BRIDGE_QT6_OVERSEER_H


#include "Worker.h"


//****************************************************************************************************************************************************
/// \brief Overseer used to manager a worker instance and its associated thread.
//****************************************************************************************************************************************************
class Overseer: public QObject
{
    Q_OBJECT
public: // member functions.
    explicit Overseer(Worker* worker, QObject* parent); ///< Default constructor.
    Overseer(Overseer const&) = delete; ///< Disabled copy-constructor.
    Overseer(Overseer&&) = delete; ///< Disabled assignment copy-constructor.
    ~Overseer() override; ///< Destructor.
    Overseer& operator=(Overseer const&) = delete; ///< Disabled assignment operator.
    Overseer& operator=(Overseer&&) = delete; ///< Disabled move assignment operator.
    bool isFinished() const; ///< Check if the worker is finished.

public slots:
    void startWorker(bool autorelease) const; ///< Run the worker.
    void release(); ///< Delete the worker and its thread.

public: // data members.
    QThread *thread_ { nullptr }; ///< The thread.
    Worker *worker_ { nullptr }; ///< The worker.
};


typedef std::unique_ptr<Overseer> UPOverseer; ///< Type definition for unique pointer to Overseer.
typedef std::shared_ptr<Overseer> SPOverseer; ///< Type definition for shared pointer to Overseer.


#endif //BRIDGE_QT6_OVERSEER_H

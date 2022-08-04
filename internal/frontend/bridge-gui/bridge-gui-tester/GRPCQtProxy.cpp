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


#include "GRPCQtProxy.h"
#include "MainWindow.h"


//****************************************************************************************************************************************************
//
//****************************************************************************************************************************************************
GRPCQtProxy::GRPCQtProxy()
    : QObject(nullptr)
{
}

//****************************************************************************************************************************************************
//
//****************************************************************************************************************************************************
void GRPCQtProxy::connectSignals()
{
    MainWindow &mainWindow = app().mainWindow();
    SettingsTab &settingsTab = mainWindow.settingsTab();
    UsersTab &usersTab = mainWindow.usersTab();
    connect(this, &GRPCQtProxy::delayedEventRequested, &mainWindow, &MainWindow::sendDelayedEvent);
    connect(this, &GRPCQtProxy::setIsAutostartOnReceived, &settingsTab, &SettingsTab::setIsAutostartOn);
    connect(this, &GRPCQtProxy::setIsBetaEnabledReceived, &settingsTab, &SettingsTab::setIsBetaEnabled);
    connect(this, &GRPCQtProxy::setColorSchemeNameReceived, &settingsTab, &SettingsTab::setColorSchemeName);
    connect(this, &GRPCQtProxy::reportBugReceived, &settingsTab, &SettingsTab::setBugReport);
    connect(this, &GRPCQtProxy::setIsStreamingReceived, &settingsTab, &SettingsTab::setIsStreaming);
    connect(this, &GRPCQtProxy::setClientPlatformReceived, &settingsTab, &SettingsTab::setClientPlatform);
    connect(this, &GRPCQtProxy::changePortsReceived, &settingsTab, &SettingsTab::changePorts);
    connect(this, &GRPCQtProxy::setUseSSLForSMTPReceived, &settingsTab, &SettingsTab::setUseSSLForSMTP);
    connect(this, &GRPCQtProxy::setIsDoHEnabledReceived, &settingsTab, &SettingsTab::setIsDoHEnabled);
    connect(this, &GRPCQtProxy::changeLocalCacheReceived, &settingsTab, &SettingsTab::changeLocalCache);
    connect(this, &GRPCQtProxy::setIsAutomaticUpdateOnReceived, &settingsTab, &SettingsTab::setIsAutomaticUpdateOn);
    connect(this, &GRPCQtProxy::setUserSplitModeReceived, &usersTab, &UsersTab::setUserSplitMode);
    connect(this, &GRPCQtProxy::removeUserReceived, &usersTab, &UsersTab::removeUser);
    connect(this, &GRPCQtProxy::logoutUserReceived, &usersTab, &UsersTab::logoutUser);
    connect(this, &GRPCQtProxy::setUserSplitModeReceived, &usersTab, &UsersTab::setUserSplitMode);
    connect(this, &GRPCQtProxy::configureUserAppleMailReceived, &usersTab, &UsersTab::configureUserAppleMail);
}


//****************************************************************************************************************************************************
/// \param[in] event The event.
//****************************************************************************************************************************************************
void GRPCQtProxy::sendDelayedEvent(bridgepp::SPStreamEvent const &event)
{
    emit delayedEventRequested(event);
}


//****************************************************************************************************************************************************
/// \param[in] on The value.
//****************************************************************************************************************************************************
void GRPCQtProxy::setIsAutostartOn(bool on)
{
    emit setIsAutostartOnReceived(on);
}


//****************************************************************************************************************************************************
/// \param[in] enabled The value.
//****************************************************************************************************************************************************
void GRPCQtProxy::setIsBetaEnabled(bool enabled)
{
    emit setIsBetaEnabledReceived(enabled);
}


//****************************************************************************************************************************************************
/// \param[in] name The color scheme.
//****************************************************************************************************************************************************
void GRPCQtProxy::setColorSchemeName(QString const &name)
{
    emit setColorSchemeNameReceived(name);
}


//****************************************************************************************************************************************************
/// \param[in] osType The OS type.
/// \param[in] osVersion The OS version.
/// \param[in] emailClient The email client.
/// \param[in] address The address.
/// \param[in] description The description.
/// \param[in] includeLogs Should the logs be included.
//****************************************************************************************************************************************************
void GRPCQtProxy::reportBug(QString const &osType, QString const &osVersion, QString const &emailClient, QString const &address,
    QString const &description, bool includeLogs)
{
    emit reportBugReceived(osType, osVersion, emailClient, address, description, includeLogs);
}


//****************************************************************************************************************************************************
/// \param[in] isStreaming Is the gRPC server streaming.
//****************************************************************************************************************************************************
void GRPCQtProxy::setIsStreaming(bool isStreaming)
{
    emit setIsStreamingReceived(isStreaming);
}


//****************************************************************************************************************************************************
/// \param[in] clientPlatform The client platform.
//****************************************************************************************************************************************************
void GRPCQtProxy::setClientPlatform(QString const &clientPlatform)
{
    emit setClientPlatformReceived(clientPlatform);
}


//****************************************************************************************************************************************************
/// \param[in] imapPort The IMAP port
/// \param[in] smtpPort The SMTP port
//****************************************************************************************************************************************************
void GRPCQtProxy::changePorts(qint32 imapPort, qint32 smtpPort)
{
    emit changePortsReceived(imapPort, smtpPort);
}


//****************************************************************************************************************************************************
/// \param[in] use Should SMTP use SSL?
//****************************************************************************************************************************************************
void GRPCQtProxy::setUseSSLForSMTP(bool use)
{
    emit setUseSSLForSMTPReceived(use);
}


//****************************************************************************************************************************************************
/// \param[in] enabled Is DoH enabled?
//****************************************************************************************************************************************************
void GRPCQtProxy::setIsDoHEnabled(bool enabled)
{
    emit setIsDoHEnabledReceived(enabled);
}


//****************************************************************************************************************************************************
/// \param[in] enabled is cache on disk enabled?
/// \param[in] path The path for the cache on disk.
//****************************************************************************************************************************************************
void GRPCQtProxy::changeLocalCache(bool enabled, QString const &path)
{
    emit changeLocalCacheReceived(enabled, path);
}


//****************************************************************************************************************************************************
/// \param[in] on Is automatic update on?
//****************************************************************************************************************************************************
void GRPCQtProxy::setIsAutomaticUpdateOn(bool on)
{
    emit setIsAutomaticUpdateOnReceived(on);
}


//****************************************************************************************************************************************************
/// \param[in] userID The userID.
/// \param[in] makeItActive Should split mode be active.
//****************************************************************************************************************************************************
void GRPCQtProxy::setUserSplitMode(QString const &userID, bool makeItActive)
{
    emit setUserSplitModeReceived(userID, makeItActive);
}


//****************************************************************************************************************************************************
/// \param[in] userID The userID.
//****************************************************************************************************************************************************
void GRPCQtProxy::logoutUser(QString const &userID)
{
    emit logoutUserReceived(userID);
}


//****************************************************************************************************************************************************
/// \param[in] userID The userID.
//****************************************************************************************************************************************************
void GRPCQtProxy::removeUser(QString const &userID)
{
    emit removeUserReceived(userID);
}


//****************************************************************************************************************************************************
/// \param[in] userID The userID.
/// \param[in] address The address.
//****************************************************************************************************************************************************
void GRPCQtProxy::configureUserAppleMail(QString const &userID, QString const &address)
{
    emit configureUserAppleMailReceived(userID, address);
}
